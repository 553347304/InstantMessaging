import {ElMessage} from "element-plus";

class m_message {
    info = (s: string) => ElMessage.info(s);
    success = (s: string) => ElMessage.success(s);
    warning = (s: string) => ElMessage.warning(s);
    error = (s: string) => ElMessage.error(s);
}

// export interface Message extends MessageFn {
//     success: MessageTypedFn;
//     warning: MessageTypedFn;
//     info: MessageTypedFn;
//     error: MessageTypedFn;
// }
export const Message = new m_message();