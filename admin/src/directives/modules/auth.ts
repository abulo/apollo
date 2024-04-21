/**
 * v-auth
 * 按钮权限指令
 */
import { Directive, DirectiveBinding } from "vue";
import { HasPermissionItem } from "@/utils/permission";

const auth: Directive = {
  mounted(el: HTMLElement, binding: DirectiveBinding) {
    const { value } = binding;
    if (value instanceof Array && value.length) {
      let hasPermission = false;
      for (let index = 0; index < value.length; index++) {
        hasPermission = HasPermissionItem(value[index]);
        if (hasPermission) break;
      }
      if (!hasPermission) el.remove();
    } else {
      if (!HasPermissionItem(value)) el.remove();
    }
  }
};

export default auth;
