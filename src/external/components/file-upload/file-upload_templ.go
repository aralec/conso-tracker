// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package fileupload

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func FileUpload() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"file-upload-component\" class=\"file has-name\"><label class=\"file-label\"><input id=\"file-input\" class=\"file-input\" type=\"file\" name=\"file\"> <span class=\"file-cta\"><span class=\"file-icon\"><i class=\"iconoir-upload\"></i></span> <span class=\"file-label\">Sélectionnez un fichier</span></span> <span class=\"file-name\">Aucun fichier importé</span></label></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = initFileUploader().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func initFileUploader() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_initFileUploader_bdb5`,
		Function: `function __templ_initFileUploader_bdb5(){const uploader = htmx.find('#file-upload-component');
    const fileInput = htmx.find('#file-input');
    fileInput.onchange = (event) => {
        console.log('File uploaded', event);
        console.log('File input files', fileInput.files);
        if (fileInput.files.length > 0) {
            const fileName = htmx.find(uploader, '.file-name');
            fileName.textContent = fileInput.files[0].name;
        }
    };
}`,
		Call:       templ.SafeScript(`__templ_initFileUploader_bdb5`),
		CallInline: templ.SafeScriptInline(`__templ_initFileUploader_bdb5`),
	}
}
