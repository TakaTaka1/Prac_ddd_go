package application

import 

type TaskApplication struct {
	ctx context.Context
	taskRepo repository.TaskRepository
}

func (s *TaskApplication) CreateTask(name string, dueDate time.Time, priority domain.PriorityStatus) error {
	task, err := domain.NewTask(name, dueDate, priority)
	if err != nil {
		return err
	}
	if err := s.taskRepo.Save(ctx, task); err != nil {
		return err
	}
	return nil
}

func (s *TaskApplication) PostponeTask(taskID string) error {
	task, err := s.taskRepo.GetByID(ctx, taskID)
	if err != nil {
		return err
	}

	postponedTask, err := task.Postpone()
	if err != nil {
		return err
	}
	if err := s.taskRepo.Save(ctx, postponedTask); err != nil {
		return err
	}
	return nil	
}


