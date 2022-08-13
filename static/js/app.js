

      function Prompt() {
        let toast = function (c) {
          const { icon = 'success', position = 'top-end', title = '' } = c
          const Toast = Swal.mixin({
            toast: true,
            position,
            title: msg,
            icon,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
              toast.addEventListener('mouseenter', Swal.stopTimer)
              toast.addEventListener('mouseleave', Swal.resumeTimer)
            },
          })

          Toast.fire()
        }
        let success = function (c) {
          const { footer = '', msg = '', title = '' } = c
          Swal.fire({
            icon: 'success',
            title: title,
            text: msg,
            footer,
          })
        }
        let error = function (c) {
          const { footer = '', msg = '', title = '' } = c
          Swal.fire({
            icon: 'error',
            title: title,
            text: msg,
            footer,
          })
        }
        async function custom(c) {
          const { msg = '', title = '',icon="",showConfirmButton=true } = c
          const { value: result } = await Swal.fire({
            icon,
            title: title,
            html: msg,
            backdrop: false,
            focusConfirm: false,
            showCancelButton: true,
            showConfirmButton,
            willOpen: () => {
              if (c.willOpen !== undefined) {
                c.willOpen()
              }
            },
            didOpen: () => {
              if (c.didOpen !== undefined) {
                c.didOpen()
              }
            },
          })
          if (result) {
            if (result.dismiss !== Swal.DismissReason.cancel) {
              if (result.value !== '') {
                if (c.callback !== undefined) {
                  c.callback(result)
                }
              } else {
                c.callback(false)
              }
            } else {
              c.callback(false)
            }
          }
        }
        return { toast, success, error, custom }
      }