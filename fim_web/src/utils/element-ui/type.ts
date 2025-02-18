export namespace m_type_el {
    export namespace props {
        export interface AvatarCropper {
            type: string // 上传类型, 企业logo / 浏览器logo
            allowTypeList: string[] // 接收允许上传的图片类型
            limitSize: number // 限制大小
            fixedNumber: number[] // 截图框的宽高比例
            fixedNumberAider?: number[] // 侧边栏收起截图框的宽高比例
            previewWidth: number // 预览宽度
            title?: string // 裁剪标题
        }
        export const AvatarCropper = (): AvatarCropper => ({
            type: 'systemLogo',
            allowTypeList:  ['jpg', 'png', 'jpeg'],
            limitSize: 1,
            fixedNumber:  [1, 1],
            fixedNumberAider: [1, 1],
            previewWidth: 228,
            title: 'LOGO裁剪'
        })
    }

    export namespace form {
        export interface option {
            value: string | number
            label: string
        }

        export interface Input {
            label?: string
            value: string
            type?: "text" | "textarea" | "password" | "button" | "checkbox" | "file" | "number" | "radio"
            disabled?: boolean  // 是否禁用
            max?: number
        }

        export interface Select {
            label: string
            value: string | number
            option: option[]
        }

        export interface Switch {
            label: string
            value: boolean
        }

        export interface Avatar {
            label: string
            src: string
        }
    }


}
