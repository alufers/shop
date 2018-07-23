<template>
    <div>
        <v-breadcrumbs>
            <v-icon slot="divider">mdi-chevron-right</v-icon>
            <v-breadcrumbs-item nuxt to="/admin" exact>Dashboard</v-breadcrumbs-item>
            <v-breadcrumbs-item nuxt to="/admin/products" exact>Products</v-breadcrumbs-item>
        </v-breadcrumbs>
        <div class="elevation-1">
            <v-toolbar flat color="white">
                <v-toolbar-title>Products</v-toolbar-title>
                <v-spacer></v-spacer>

            </v-toolbar>
        </div>
    </div>
</template>

<script>
export default {
  layout: "admin"
};
</script>

<style>
</style>
