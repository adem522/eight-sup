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
      <v-col cols="6"><!-- item-text will change-->
        <v-combobox
          style="width: 100%"
          v-model="selectedSeller"
          @change="takeSellerUsername()"
          :items="wantedUsers"
          label="Buyer Username"
          item-text="buyerUsername" 
          item-value="buyerUsername"
          clearable small-chips outlined dense
        />
      </v-col>
      <v-col cols="6"><!-- item-text will change-->
        <v-combobox
          style="width: 100%"
          v-model="selectedUnique"
          @change="takeUnique()"
          :items="sellerusername"
          label="Buyed Package"
          item-text="unique" 
          item-value="unique"
          clearable small-chips outlined dense
        />
      </v-col>
       <v-col cols="6">
         Completed
        <v-card
          class="mx-auto"
          max-width="800"
          tile
          v-for="item in unique" :key="item.prop"
          >
          <v-list-item three-line v-if="item.status=='completed'">
              <v-list-item-content>
                  <v-list-item-title>{{item.prop}}</v-list-item-title>
                  <v-list-item-action> 
                    <v-btn @click="wantProp(item.unique,item.prop,'')" >Want</v-btn>
                  </v-list-item-action>
              </v-list-item-content>
          </v-list-item>
        </v-card>
      </v-col>
       <v-col cols="6" >
         Requested
        <v-card
          class="mx-auto"
          max-width="800"
          tile
          v-for="item in unique" :key="item.prop"
          >
          <v-list-item three-line v-if="item.status=='requested'">
              <v-list-item-content>
                  <v-list-item-title>{{item.prop}}</v-list-item-title>                  
                  <v-list-item-action> 
                    <v-btn @click="wantProp(item.unique,item.prop,'complete')" >Complete</v-btn>
                  </v-list-item-action>
              </v-list-item-content>
          </v-list-item>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from "axios";
import { mapState } from "vuex";
export default {
  name: "WantPropStreamer",
  computed: {
    ...mapState(["username", "host","type"]),
  },
  data() {
    return {
      sellerusername:[],
      storedPlan:[],
      wantedUsers:[],
      unique:[],
      selectedSeller: null,
      selectedUnique: null,
    };
  },
  methods: {
    wantProp(uniqueCame,propCame,statusCame){
      axios
      .post(this.host+"/create/wantStreamer",{
        unique:uniqueCame,
        prop:propCame,
        buyerUsername:this.selectedSeller,
        sellerUsername:this.username,
        status:statusCame,
        output:"json"
      })
      .catch((error) => {
          window.alert(`The API returned an error: ${error.response.data}`);
      });
    },
    takeSellerUsername() {
      this.sellerusername=[]//for delete datas in combobox
      this.unique=[]        //for delete datas in combobox
      for(let prop in this.storedPlan){
        for(let prop2 in this.storedPlan[prop].package.items){
            if (this.selectedSeller == this.storedPlan[prop].package.items[prop2].buyerUsername){
              this.sellerusername.push(this.storedPlan[prop].package)
          }
        }
      }
    },
    takeWantedUsers(){
      for(let prop in this.storedPlan){
        for(let prop2 in this.storedPlan[prop].package.items){
          this.wantedUsers.push(this.storedPlan[prop].package.items[prop2].buyerUsername)
        }
      }
    },
    takeUnique() {
      this.unique=[]
      for(let prop in this.storedPlan){
        if (this.selectedUnique.unique == this.storedPlan[prop].package.unique){
            for(let prop2 in this.storedPlan[prop].package.items){
              if (this.selectedSeller == this.storedPlan[prop].package.items[prop2].buyerUsername){
              this.unique.push({
                prop:this.storedPlan[prop].package.items[prop2].prop,
                status:this.storedPlan[prop].package.items[prop2].status,
                unique:this.storedPlan[prop].package.unique,
              })
            }
          }
        }
      }
    },
    getStocks() {
      axios
        .put(this.host + "/list/userPlan", {
          username: this.username,
        }
      )
        .then((response) => {
          this.storedPlan=response.data[0].plan;
          this.takeWantedUsers();
        }
      );
    },
  },
  created() {
      this.getStocks();
  },
};
</script>
