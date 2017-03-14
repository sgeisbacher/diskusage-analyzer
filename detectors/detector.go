package detectors

type HotspotsDetector interface {
	Detect(ctx AnalyzerContext, top int) (Dirs, error)
}
