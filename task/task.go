package task

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type TaskStatus string

const (
	TaskStatusTodo       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in-progress"
	TaskStatusDone       TaskStatus = "done"
)

type Task struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	Status    TaskStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt time.Time  `json:"deleted_at"`
}

func LoadTasks() ([]*Task, error) {
	jsonFile, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening tasks.json:", err)
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	if len(byteValue) == 0 {
		return []*Task{}, nil
	}

	var tasks []*Task
	if err = json.Unmarshal(byteValue, &tasks); err != nil {
		fmt.Println("Error unmarshalling tasks.json:", err)
		return nil, err
	}
	return tasks, nil
}

func SaveTasks(tasks []*Task) error {
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", jsonData, 0644)
}

func AddTask(name string) error {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return err
	}

	tasks = append(tasks, &Task{
		Id:        fmt.Sprint(len(tasks) + 1),
		Name:      name,
		Status:    TaskStatusTodo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err = SaveTasks(tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		return err
	}

	fmt.Printf("Task added successfully (ID: %s)\n", tasks[len(tasks)-1].Id)
	return nil
}

func UpdateTask(id, name string) error {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return err
	}

	for _, task := range tasks {
		if task.Id == id {
			task.Name = name
			task.UpdatedAt = time.Now()
			break
		}
	}

	if err = SaveTasks(tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		return err
	}

	return nil
}

func DeleteTask(id string) error {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return err
	}

	for _, task := range tasks {
		if task.Id == id {
			task.DeletedAt = time.Now()
			break
		}
	}

	if err = SaveTasks(tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		return err
	}

	return nil
}

// 计算字符串显示宽度，中文字符计为2，其他字符计为1
func getStrWidth(s string) int {
	width := 0
	for _, r := range s {
		if r > 0x4E00 && r < 0x9FFF { // 简单判断是否为中文字符
			width += 2
		} else {
			width += 1
		}
	}
	return width
}

func ListTasks(status string) error {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return err
	}

	maxWidth := 0
	show := make([]*Task, 0, len(tasks))
	for _, task := range tasks {
		if !task.DeletedAt.IsZero() {
			continue
		}

		if status == "" || status == string(task.Status) {
			show = append(show, task)
			maxWidth = max(maxWidth, getStrWidth(task.Name))
		}
	}

	// 打印表头
	idHeader := "ID"
	nameHeader := "Name"
	statusHeader := "Status"

	fmt.Printf("%-4s  %-*s  %s\n", idHeader, maxWidth, nameHeader, statusHeader)
	fmt.Println(strings.Repeat("-", 6+maxWidth+len(statusHeader)+2))

	for i := len(show) - 1; i >= 0; i-- {
		task := show[i]
		padding := maxWidth - getStrWidth(task.Name)
		fmt.Printf("%-4s  %s%s  %s\n",
			task.Id,
			task.Name,
			strings.Repeat(" ", padding),
			task.Status,
		)
	}

	return nil
}

func MarkTaskInProgress(id string) error {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return err
	}

	for _, task := range tasks {
		if task.Id == id {
			task.Status = TaskStatusInProgress
			task.UpdatedAt = time.Now()
			break
		}
	}

	if err = SaveTasks(tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		return err
	}

	return nil
}

func MarkTaskDone(id string) error {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return err
	}

	for _, task := range tasks {
		if task.Id == id {
			task.Status = TaskStatusDone
			task.UpdatedAt = time.Now()
			break
		}
	}

	if err = SaveTasks(tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		return err
	}

	return nil
}
