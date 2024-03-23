import { useAuthStore } from "@/stores/modules/auth";

export const HasPermission = (...permissionList: string[]) => {
  if (permissionList.length === 1) {
    let userHasPermission = true;
    let permissionItem = permissionList[0];
    if (!HasPermissionItem(permissionItem)) {
      userHasPermission = false;
    }
    return userHasPermission;
  } else {
    let userHasPermission = false;
    for (let i = 0; i < permissionList.length; i++) {
      let permissionItem = permissionList[i];
      if (HasPermissionItem(permissionItem)) {
        userHasPermission = true;
        break;
      }
    }
    return userHasPermission;
  }
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
