<template>
  <div style="width: 104%">
    <v-container fluid>
      <v-row>
        <v-col cols="12">
          <v-combobox
            v-model="selectedItem"
            :items="items"
            :label="label"
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
//import axios from 'axios';
export default {
  props: {
    label: {
      type: String,
      default: "Label",
    },
    items: {
      type: Array,
      required: true,
    },
  },
  data() {
    return {
      selectedItem: "",
    };
  },
};
</script>
