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
          v-model="selectedSeller"
          :items="sellers"
          label="Sellers"
          clearable
          small-chips
          outlined
          dense
        />
      </v-col>
      <v-col cols="12">
        <v-combobox
          v-model="selectedInfo"
          :items="storedInfo"
          label="Info Names"
          item-text="name"
          item-value="unique"
          clearable
          small-chips
          outlined
          dense
        />
      </v-col>
      <v-col cols="12">
        <v-btn @click="buyPackage" style="width: 100%">BUY PACKAGE</v-btn>
      </v-col>
      <v-col v-if="success" cols="12">
        <v-alert type="success">
          Buyed Package
        </v-alert>
      </v-col>
    </v-row>
  </v-container>
</template>
<script>
import axios from "axios";
import { mapState } from "vuex";
export default {
  name: "Package",
  computed: {
    ...mapState(["username", "host","storedInfo","storedSilverStock","storedGoldStock"]),
  },
  data() {
    return {
      buyerUsername: "",
      sellerUsername: "",
      unique: "",
      sellers: [],
      selectedItems:[],
      selectedInfo: null,
      selectedSeller: null,
      success:false
    };
  },
  methods: {
    buyPackage() {
      //console.log("selectedInfo",this.selectedInfo.items,"\n")
      //console.log("items",items)
      this.selectedItems=[]
      for (let key in this.selectedInfo.items){
        //console.log("key",key,"\n")
        //console.log("data",this.selectedInfo.items[key],"\n")
        this.selectedItems.push({prop:this.selectedInfo.items[key],status:"available",buyerUsername:this.username})
      }
      ///console.log("selectedItems",this.selectedItems,"\n")
      axios
        .post(this.host + "/create/event", {
          buyerUsername: this.username,
          sellerUsername: this.selectedSeller,
          unique: this.selectedInfo.unique,
          items:this.selectedItems,
          output: "json",
        })
        .then(()=>{
          for(let prop in this.storedPlan){
              if (this.selectedInfo.unique == this.storedPlan[prop].package.unique){
              this.storedPlan[prop].package.stock = parseInt(this.take) + parseInt(this.already)
            }
          }
            this.success=true
            this.hide_alert()
        })
        .catch((error) => {
          window.alert(`The API returned an error: ${error.response.data}`);
        });
    },
    getSellers() {
      axios.get(this.host + "/list/user")
      .then((response) => {
        this.sellers = response.data.map((el) => el.username);
      });
    },
    getInfo() {
      if(this.storedInfo.length==0){
        axios
          .get(this.host + "/list/planInfo")
          .then((response) => {

            this.$store.commit("setStoredInfo", response.data.data);
        });
      }
    },
    hide_alert() {
      window.setInterval(() => {
        this.success=false;
      }, 1500)    
    }
  },
  created() {
    this.getSellers();
    this.getInfo();
  },
};
</script>
