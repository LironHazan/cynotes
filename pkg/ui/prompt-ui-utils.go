package promptUiUtils

import (
	fsutils "cynotes/pkg/fs"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/samber/lo"
	"io/fs"
	"path/filepath"
)

func PromptPasswdInput() (string, error) {
	validate := func(input string) error {
		return lo.Ternary[error](len(input) < 8, errors.New("password must have more than 8 characters"), nil)
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     "Password",
		Validate:  validate,
		Templates: templates,
		Mask:      '*',
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", nil
	}

	return result, err
}

func PromptFileName() (string, error) {
	validate := func(input string) error {

		visit := func(path string, dir fs.DirEntry, err error) error {
			if dir.Name() == input {
				fmt.Printf(" name %s", input)
				return errors.New("filename already exists")
			}
			return nil
		}

		path, _ := fsutils.GetCYNotesPath()
		err := filepath.WalkDir(path, visit)
		if err != nil {
			return errors.New("filename already exists")
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     "Filename",
		Validate:  validate,
		Templates: templates,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", nil
	}
	return result, err
}

func PromptGitClient(label string) (string, error) {
	validate := func(input string) error {
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     label,
		Validate:  validate,
		Templates: templates,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", nil
	}
	return result, err
}

func BasicPrompt(message string) (string, error) {
	prompt := promptui.Prompt{
		Label: message,
	}
	return prompt.Run()
}

func BasicPromptSelections(message string, items []string) (string, error) {
	prompt := promptui.Select{
		Label: message,
		Items: items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return result, err
}
