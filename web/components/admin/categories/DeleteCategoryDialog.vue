<template>
    <div>
        <v-dialog v-model="dialog" max-width="500px">
            <v-btn slot="activator" flat fab dark small color="red">
                <v-icon>
                    mdi-delete
                </v-icon>
            </v-btn>
            <v-card>
                <v-card-title>
                    <span class="headline">Do you really want to delete '{{category.Name}}'?</span>
                </v-card-title>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="primary darken-1" flat @click.native="close">Cancel</v-btn>
                    <v-btn color="red" dark @click.native="doDelete">Delete</v-btn>
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
      dialog: false,
      snackbar: null,
      snackbarColor: "primary"
    };
  },
  props: ["category"],
  methods: {
    close() {
      this.dialog = false;
    },
    async doDelete() {
      let url;
      try {
        await api.delete(encodeParams`/api/shop/categories/${this.category.ID}`);
        this.snackbarColor = "primary";
        this.snackbar = `Category ${this.category.Name} deleted.`;
        this.close();
      } catch (e) {
        this.snackbarColor = "error";
        this.snackbar = "Error: " + e.message;
        this.close();
      }
      this.$emit("deleted");
    }
  }
};
</script>

<style>
</style>
