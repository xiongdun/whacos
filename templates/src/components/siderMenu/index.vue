<template>
  <div class="menutree">
    <div v-for="menu in menuList" :key="menu.path">
      <el-submenu :index="menu.path" v-if="menu.children">
        <template slot="title">
          <span>{{menu.name}}</span>
        </template>
        <el-menu-item-group>
          <sider-menu :menuList="menu.children"></sider-menu>
        </el-menu-item-group>
      </el-submenu>
      <el-menu-item v-else :index="menu.path">
        <span slot="title">{{menu.name}}</span>
      </el-menu-item>
    </div>
  </div>
</template>

<script>
import SiderMenu from "components/siderMenu";
export default {
  name: "SiderMenu",
  props: {
    menuList: {
      type: Array,
      default: () => []
    }
  },
  components: {
    SiderMenu
  }
};
</script>

<style scoped>
</style>

