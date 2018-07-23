<template>
    <div>
        <v-dialog v-model="dialog" max-width="500px">
            <v-btn slot="activator" color="primary" dark class="mb-2">
                <template v-if="!parentID">New category</template>
                <template v-else>New subcategory</template>
            </v-btn>
            <v-card>
                <v-card-title>
                    <span class="headline">New category</span>
                </v-card-title>

                <v-card-text>
                    <v-container grid-list-md>
                        <v-text-field v-model="category.Name" label="Category name"></v-text-field>
                    </v-container>
                </v-card-text>

                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="primary darken-1" flat @click.native="close">Cancel</v-btn>
                    <v-btn color="primary darken-1" @click.native="save" :disabled="category.Name.trim() === ''">Save</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
        <v-snackbar :value="!!snackbar" @change="$event === false ? (snackbar = null) : null" :color="snackbarColor">
            {{ snackbar }}
            <v-btn dark flat @click="snackbar = null">
                Close
            </v-btn>
        </v-snackbar>
    </div>
</template>

<script>
import api from "~/lib/api";
import { encodeParams } from "~/lib/util";

export default {
  data() {
    return {
      product: {
        Name: ""
      },
      dialog: false,
      snackbar: null,
      snackbarColor: "primary"
    };
  },
  props: ["parentID"],
  methods: {
    close() {
      this.dialog = false;
      this.category = {
        Name: ""
      };
    },
    async save() {
     
    }
  }
};
</script>

<style>
</style>
