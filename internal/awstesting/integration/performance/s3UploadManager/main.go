// +build integration,perftest

package main

import (
	"flag"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awstesting/integration"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3manager"
)

var config Config

func main() {
	parseCommandLine()

	log.SetOutput(os.Stderr)

	var (
		file *os.File
		err  error
	)

	if config.Filename != "" {
		file, err = os.Open(config.Filename)
		if err != nil {
			log.Fatalf("failed to open file: %v", err)
		}
	} else {
		file, err = integration.CreateFileOfSize(config.TempDir, config.Size)
		if err != nil {
			log.Fatalf("failed to create file: %v", err)
		}
		defer os.Remove(file.Name())
	}

	defer file.Close()

	traces := make(chan *RequestTrace, config.SDK.Concurrency)
	requestTracer := uploadRequestTracer(traces)
	uploader := newUploader(config.Client, config.SDK, requestTracer)

	metricReportDone := make(chan struct{})
	go func() {
		defer close(metricReportDone)
		metrics := map[string]*RequestTrace{}
		for trace := range traces {
			curTrace, ok := metrics[trace.Operation]
			if !ok {
				curTrace = trace
			} else {
				curTrace.attempts = append(curTrace.attempts, trace.attempts...)
				if len(trace.errs) != 0 {
					curTrace.errs = append(curTrace.errs, trace.errs...)
				}
				curTrace.finish = trace.finish
			}

			metrics[trace.Operation] = curTrace
		}

		for _, name := range []string{
			"CreateMultipartUpload",
			"CompleteMultipartUpload",
			"UploadPart",
			"PutObject",
		} {
			if trace, ok := metrics[name]; ok {
				printAttempts(name, trace, config.LogVerbose)
			}
		}
	}()

	log.Println("starting upload...")
	start := time.Now()
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: &config.Bucket,
		Key:    aws.String(filepath.Base(file.Name())),
		Body:   file,
	})
	if err != nil {
		log.Fatalf("failed to upload object, %v", err)
	}
	close(traces)

	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	dur := time.Since(start)
	log.Printf("Upload finished, Size: %d, Dur: %s, Throughput: %.5f GB/s",
		size, dur, (float64(size)/(float64(dur)/float64(time.Second)))/float64(1e9),
	)

	<-metricReportDone
}

func parseCommandLine() {
	config.SetupFlags("", flag.CommandLine)

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		flag.CommandLine.PrintDefaults()
		log.Fatalf("failed to parse CLI commands")
	}
	if err := config.Validate(); err != nil {
		flag.CommandLine.PrintDefaults()
		log.Fatalf("invalid arguments: %v", err)
	}
}

func printAttempts(op string, trace *RequestTrace, verbose bool) {
	if !verbose {
		return
	}

	log.Printf("%s: latency:%s requests:%d errors:%d",
		op,
		trace.finish.Sub(trace.start),
		len(trace.attempts),
		len(trace.errs),
	)

	for _, a := range trace.attempts {
		log.Printf("  * %s", a)
	}
	if err := trace.Err(); err != nil {
		log.Printf("Operation Errors: %v", err)
	}
	log.Println()
}

func uploadRequestTracer(traces chan<- *RequestTrace) aws.Option {
	tracerOption := func(r *aws.Request) {
		id := "op"
		if v, ok := r.Params.(*s3.UploadPartInput); ok {
			id = strconv.FormatInt(*v.PartNumber, 10)
		}
		tracer := NewRequestTrace(r.Context(), r.Operation.Name, id)
		r.SetContext(tracer)

		r.Handlers.Send.PushFront(tracer.OnSendAttempt)
		r.Handlers.CompleteAttempt.PushBack(tracer.OnCompleteAttempt)
		r.Handlers.Complete.PushBack(tracer.OnComplete)
		r.Handlers.Complete.PushBack(func(rr *aws.Request) {
			traces <- tracer
		})
	}

	return tracerOption
}

func SetUnsignedPayload(r *aws.Request) {
	if r.Operation.Name != "UploadPart" && r.Operation.Name != "PutObject" {
		return
	}
	r.HTTPRequest.Header.Set("X-Amz-Content-Sha256", "UNSIGNED-PAYLOAD")
}

func newUploader(clientConfig ClientConfig, sdkConfig SDKConfig, options ...aws.Option) *s3manager.Uploader {
	client := NewClient(clientConfig)

	if sdkConfig.WithUnsignedPayload {
		options = append(options, SetUnsignedPayload)
	}

	cfg, err := external.LoadDefaultAWSConfig(aws.Config{HTTPClient: client})
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	svc := s3.New(cfg)
	svc.Disable100Continue = !sdkConfig.ExpectContinue

	uploader := s3manager.NewUploaderWithClient(svc, func(u *s3manager.Uploader) {
		u.PartSize = sdkConfig.PartSize
		u.Concurrency = sdkConfig.Concurrency
		u.BufferProvider = sdkConfig.BufferProvider

		u.RequestOptions = append(u.RequestOptions, options...)
	})

	return uploader
}