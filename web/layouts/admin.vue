<template>
  <v-app>
    <v-navigation-drawer class="grey lighten-4" app>
      <v-list dense class="grey lighten-4">
        <v-list-tile nuxt to="/admin/categories">
          <v-list-tile-action>
            <v-icon>mdi-tag</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title class="grey--text">
              Categories
            </v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
        <v-list-tile nuxt to="/admin/products">
          <v-list-tile-action>
            <v-icon>mdi-tshirt-crew</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>
            <v-list-tile-title class="grey--text">
              Products
            </v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>
    <v-content>
      <v-container fluid>
        <nuxt/>
      </v-container>
    </v-content>
  </v-app>
</template>
<style>
</style>
<script>
</script>

