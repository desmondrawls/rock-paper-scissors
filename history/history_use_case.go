package history

import "github.com/desmondrawls/rock-paper-scissors/models"

type UseCase struct {
    Repository Repository
}

//go:generate counterfeiter --fake-name VisualizerSpy . Visualizer
type Visualizer interface {
    Empty()
    Error()
    Records([]models.Record)
}

//go:generate counterfeiter --fake-name RepositoryStub . Repository
type Repository interface {
    List() ([]models.Record, error)
}

func (u *UseCase) Execute(visualizer Visualizer) {
    records, err := u.Repository.List()
    if err != nil {
        visualizer.Error()
        return
    }
    if len(records) == 0 {
        visualizer.Empty()
        return
    }
    visualizer.Records(records)
}
