<template>
  <div style="width: 104%">
    <v-container fluid>
      <v-row>
        <v-col cols="12">
          <v-combobox
            v-model="select"
            :items="items"
            label="Info Names"
            clearable
            small-chips
            outlined
            dense
          />
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import axios from "axios";
import { mapState } from "vuex";
export default {
  computed: {
    ...mapState(["host"]),
  },
  data() {
    return {
      select: ["Silver"],
      items: ["Silver", "Gold"],
    };
  },
  created() {
    var user = this;
    axios.get(this.host + "/list/user").then(function (response) {
      user.listDataString = JSON.stringify(response.data, "username", "\t");
      //user.items=response.data
    });
  },
};
</script>
