package build

// StageType is a build stage type
type StageType string

const (
	// CREATED StageType
	CREATED = StageType("created")
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
	case CREATED:
		return nil
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
	case CREATED:
		if predecessor == nil {
			return nil
		}
	case STARTED:
		if predecessor != nil && predecessor.Type == CREATED {
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
