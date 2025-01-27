package company

type Worker interface {
	Work(tasks []string) string
}

type Company struct {
	personal []Worker
}

func (c *Company) Hire(newbie Worker) {
	c.personal = append(c.personal, newbie)
}

func (c Company) Process(id int, tasks []string) (res string) {
	return c.personal[id].Work(tasks)
}
