package promptUiUtils

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
)

func PromptPasswdInput() (string, error) {
	validate := func(input string) error {
		if len(input) < 8 {
			// Todo: add passwd pattern validation
			return errors.New("password must have more than 8 characters")
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

	fmt.Printf("Your password is %q\n", result)
	return result, err
}
