package tasks

import "context"

func (s *serv) DeleteTaskById(ctx context.Context, id int) error {
	err := s.txManger.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.tasksRepository.DeleteTaskLabelById(ctx, id)
		if errTx != nil {
			return errTx
		}

		errTx = s.tasksRepository.DeleteTaskById(ctx, id)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
