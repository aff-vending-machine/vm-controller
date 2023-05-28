package frontend

import "context"

func (w *wsImpl) SendToIdle(ctx context.Context) error {
	if err := checkConnection(w.client); err != nil {
		return err
	}

	w.mu.Lock()
	defer w.mu.Unlock()
	payload := PayloadModel{
		Code:  200,
		Stage: "idle",
	}

	return w.client.WriteJSON(payload)
}
