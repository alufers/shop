<template>
    <div>
        <v-breadcrumbs>
            <v-icon slot="divider">mdi-chevron-right</v-icon>
            <v-breadcrumbs-item nuxt to="/admin" exact>Dashboard</v-breadcrumbs-item>
            <v-breadcrumbs-item nuxt to="/admin/categories" exact>Categories</v-breadcrumbs-item>
            <v-breadcrumbs-item v-for="cb in categoryBreadcrumbs" :key="cb.ID" nuxt :to="{name: 'category-with-parent', params: {parentCategory: cb.ID}}" exact>
                {{cb.Name}}
            </v-breadcrumbs-item>
        </v-breadcrumbs>
        <div class="elevation-1">
            <v-toolbar flat color="white">
                <v-toolbar-title>Categories</v-toolbar-title>
                <v-spacer></v-spacer>
                <NewCategoryDialog :parentID="parentCategory" @saved="loadData" />
            </v-toolbar>
            <v-data-table :headers="headers" :items="items" :loading="loading" :rows-per-page-items="[25, 50, -1]">
                <template slot="items" slot-scope="props">
                    <tr @click="$router.push({name: 'category-with-parent', params: {parentCategory: props.item.ID}})">
                        <td>{{ props.item.Name }}</td>

                        <td class="justify-center layout px-0">
                            <v-badge overlap :value="props.item.ChildrenCount > 0">
                                <span slot="badge">{{props.item.ChildrenCount}}</span>
                                <v-btn flat fab dark small color="primary" nuxt :to="{name: 'category-with-parent', params: {parentCategory: props.item.ID}}">
                                    <v-icon>
                                        mdi-format-list-bulleted
                                    </v-icon>
                                </v-btn>
                            </v-badge>
                            <v-btn flat fab dark small color="primary">
                                <v-icon>
                                    mdi-pencil
                                </v-icon>
                            </v-btn>
                            <DeleteCategoryDialog :category="props.item" @deleted="loadData"/>
                        </td>
                    </tr>
                </template>
            </v-data-table>
        </div>
    </div>
</template>

<script>
import api from "~/lib/api";
import { encodeParams } from "~/lib/util";
import NewCategoryDialog from "~/components/admin/categories/NewCategoryDialog";
import DeleteCategoryDialog from "~/components/admin/categories/DeleteCategoryDialog";

const headers = [
  {
    text: "Name",
    sortable: false,
    value: "Name"
  },
  {
    text: "Actions",
    sortable: false,
    value: "Name",
    width: 10
  }
];

export default {
  layout: "admin",
  components: {
    NewCategoryDialog,
    DeleteCategoryDialog
  },
  asyncData({ params }) {
    return {
      parentCategory: params.parentCategory || null
    };
  },
  data() {
    return {
      headers,
      loading: false,
      error: null,
      items: [],
      parentCategoryInfo: null
    };
  },
  methods: {
    async loadData() {
      console.log("Loading categories...");
      this.loading = true;
      let url;
      if (this.parentCategory) {
        url = encodeParams`/api/shop/categories/${
          this.parentCategory
        }/children`;
      } else {
        url = "/api/shop/categories";
      }
      try {
        this.items = (await api.get(url)).data;
        if (this.parentCategory) {
          this.parentCategoryInfo = (await api.get(
            `/api/shop/categories/${this.parentCategory}`
          )).data;
        }
      } catch (e) {
        console.error(e);
        this.error = e.message;
      }
      this.loading = false;
    }
  },
  created() {
    this.loadData();
  },
  computed: {
    categoryBreadcrumbs() {
      if (!this.parentCategoryInfo) {
        return null;
      }
      const append = (arr, cat) =>
        cat.Parent ? (arr.unshift(cat.Parent), append(arr, cat.Parent)) : arr;
      return append([this.parentCategoryInfo], this.parentCategoryInfo);
    }
  }
};
</script>

<style scoped>
tr {
  cursor: pointer;
}
</style>
