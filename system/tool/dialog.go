package tool

import (
	"mangav4/system/app"
	"mangav4/system/app/helper"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Dialog struct {
	Status string
}

func NewDialog() *Dialog {
	return &Dialog{}
}

func (f *Dialog) MessageBox(msg string, title string) (string, error) {

	return runtime.MessageDialog(*app.WailsContext, runtime.MessageDialogOptions{
		Title:   title,
		Message: msg,
		Type:    runtime.InfoDialog,
	})
}

func (f *Dialog) MessageBoxError(msg string, title string) (string, error) {

	return runtime.MessageDialog(*app.WailsContext, runtime.MessageDialogOptions{
		Title:   title,
		Message: msg,
		Type:    runtime.ErrorDialog,
	})
}

func (f *Dialog) MessageBoxQuestion(msg string, title string) (string, error) {

	return runtime.MessageDialog(*app.WailsContext, runtime.MessageDialogOptions{
		Title:   title,
		Message: msg,
		Type:    runtime.QuestionDialog,
	})
}

/*
SaveDialog
*/
func (f *Dialog) SaveDialog(title, filename, directory string, filters []map[string]string) (string, error) {

	var filter []runtime.FileFilter
	for _, f := range filters {

		filter = append(filter, runtime.FileFilter{
			DisplayName: f["name"],
			Pattern:     f["pattern"],
		})
	}

	return runtime.SaveFileDialog(*app.WailsContext, runtime.SaveDialogOptions{
		Title:            title,
		DefaultDirectory: directory,
		DefaultFilename:  filename,
		Filters:          filter,
	})
}

/*
OpenDialog
*/
func (f *Dialog) OpenDialog(title, filename, directory string, filters []map[string]string) (string, error) {

	var filter []runtime.FileFilter
	for _, f := range filters {

		filter = append(filter, runtime.FileFilter{
			DisplayName: f["name"],
			Pattern:     f["pattern"],
		})
	}

	return runtime.OpenDirectoryDialog(*app.WailsContext, runtime.OpenDialogOptions{
		Title:            title,
		DefaultDirectory: directory,
		DefaultFilename:  filename,
		Filters:          filter,
	})
}

func (f *Dialog) OpenExpoler(path string) error {
	err := helper.OpenExpoler(path)
	if err != nil {
		return err
	}
	return nil
}
