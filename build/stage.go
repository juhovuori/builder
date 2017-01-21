package build

import "time"

// StageType is a build stage type
type StageType string

const (
	// STARTED StageType
	STARTED = StageType("started")
	// PROGRESS StageType
	PROGRESS = StageType("progress")
	// ABORTED StageType
	ABORTED = StageType("aborted")
	// FAILURE StageType
	FAILURE = StageType("failure")
	// SUCCESS StageType
	SUCCESS = StageType("success")
)

// Validate returns true for a valid StageType
func (t StageType) Validate() error {
	switch t {
	case STARTED:
		return nil
	case PROGRESS:
		return nil
	case ABORTED:
		return nil
	case FAILURE:
		return nil
	case SUCCESS:
		return nil
	}
	return ErrStageType
}

// Stage is a build stage
type Stage struct {
	Type      StageType `json:"type"`
	Timestamp int64     `json:"timestamp"`
	Name      string    `json:"name"`
	Data      []byte
}

// ValidateWithPredecessor validates a stage together with its predecessor
func (s Stage) ValidateWithPredecessor(predecessor *Stage) error {
	err := s.Type.Validate()
	if err != nil {
		return err
	}
	if predecessor != nil && s.Timestamp < predecessor.Timestamp {
		return ErrStageOrder
	}
	switch s.Type {
	case STARTED:
		if predecessor == nil {
			return nil
		}
	case PROGRESS:
		fallthrough
	case ABORTED:
		fallthrough
	case FAILURE:
		fallthrough
	case SUCCESS:
		if predecessor != nil && (predecessor.Type == STARTED || predecessor.Type == PROGRESS) {
			return nil
		}
	}
	return ErrStageOrder
}

// StartStage creates a build start stage
func StartStage() Stage {
	return Stage{
		Type:      STARTED,
		Name:      "started",
		Timestamp: time.Now().UnixNano(),
	}
}

// SuccessStage creates a build end stage
func SuccessStage() Stage {
	return Stage{
		Type:      FAILURE,
		Name:      "end-of-script",
		Timestamp: time.Now().UnixNano(),
	}
}

// FailureStage creates a build end stage
func FailureStage() Stage {
	return Stage{
		Type:      FAILURE,
		Name:      "end-of-script",
		Timestamp: time.Now().UnixNano(),
	}
}
