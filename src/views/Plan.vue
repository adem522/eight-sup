<template>
  <v-container
    fluid
    style="
      width: 50%;
      display: flex;
      align-items: center;
      justify-content: flex-start;
      flex-direction: column;
      margin: 0 auto;
    "
  >
    <v-row>
      <v-col cols="12">
        <v-combobox
          style="width: 100%"
          v-model="selectedInfo"
          @change="changeNumber()"
          :items="info"
          label="Package Names"
          item-text="name"
          item-value="unique"
          clearable
          small-chips
          outlined
          dense
        />
      </v-col>
      <v-col cols="5.5">
        <v-text-field
          v-model="already"
          disabled
          label="Already Have"
          style="width: 100%"
        />
      </v-col>
      <v-icon>mdi-plus</v-icon>
      <v-col cols="5.5">
        <v-text-field
          v-model="take"
          :rules="[rules.required]"
          label="Number You Take Stock"
          style="width: 100%"
        />
      </v-col>
      <v-col cols="12">
        <v-btn @click="pushPlan" style="width: 100%">BUY PACKAGE</v-btn>
      </v-col>
      <v-col v-if="success" cols="12">
        <v-alert type="success"> Added Package </v-alert>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from "axios";
import { mapState } from "vuex";
export default {
  name: "Plan",
  computed: {
    ...mapState([
      "username",
      "host",
      "type",
      "storedInfo"
    ]),
  },
  data() {
    return {
      buyerUsername: "",
      already: 0,
      take: 0,
      unique: "",
      info: [null],
      selectedInfo: null,
      goldStock: 0,
      success: false,
      storedPlan:[],
      rules: {
        required: (value) => !!value || "Required.",
        min: (v) => v.length >= 1 || "Min 1 characters", //issues#4
      },
    };
  },
  methods: {
    pushPlan() {
      axios
        .post(this.host + "/plan/push", {
          username: this.username,
          unique: this.selectedInfo.unique,
          number: parseInt(this.take),
          output: "json",
        })
        .then(()=>{
        for(let prop in this.storedPlan){
            if (this.selectedInfo.unique == this.storedPlan[prop].package.unique){
            this.storedPlan[prop].package.stock = 
                    parseInt(this.take) + 
                    parseInt(this.already)
          }
        }
          this.success = true;
          this.hide_alert();
        })
        .catch((error) => {
          window.alert(`The API returned an error: ${error}`);
        });
    },
   changeNumber() {
     this.already=0
      for(let prop in this.storedPlan){
          if (this.selectedInfo.unique == this.storedPlan[prop].package.unique){
          this.already = this.storedPlan[prop].package.stock
        }
      }
    },
    getStocks() {
      axios
        .put(this.host + "/list/userPlan", {
          username: this.username,
        })
        .then((response) => {
          this.storedPlan=response.data[0].plan;
        });
    },
    getInfo() {
      if (this.storedInfo.length == 0) {
        axios.get(this.host + "/list/planInfo").then((response) => {
          this.info = response.data.data;
          this.$store.commit("setStoredInfo", response.data.data);
        });
      } else {
        this.info = this.storedInfo;
      }
    },
    hide_alert() {
      window.setInterval(() => {
        this.success = false;
      }, 1500);
    },
  },
  created() {
    this.getStocks();
    this.getInfo();
  },
};
</script>
