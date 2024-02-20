import { useAuthStore } from "@/stores/modules/auth";

export const HasManyPermission = (...permissionList: string[]) => {
  let userHasPermission = false;
  for (let i = 0; i < permissionList.length; i++) {
    let permissionItem = permissionList[i];
    if (HasPermissionItem(permissionItem)) {
      userHasPermission = true;
      break;
    }
  }
  return userHasPermission;
};

export const HasPermission = (...permissionList: string[]) => {
  // 判断是permission字符串是否存在 : 符号
  let userHasPermission = true;
  for (let i = 0; i < permissionList.length; i++) {
    let permissionItem = permissionList[i];
    if (!HasPermissionItem(permissionItem)) {
      userHasPermission = false;
      break;
    }
  }
  return userHasPermission;
};

export const HasPermissionItem = (permission: string) => {
  const authStore = useAuthStore();
  let itemName = "";
  if (permission.indexOf(".") > -1) {
    const permissionArr = permission.split(".");
    itemName = permissionArr[0];
  } else {
    itemName = permission;
  }
  const authButtons = authStore.authButtonListGet[itemName] || [];
  return authButtons.includes(permission);
};
