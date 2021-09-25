<template>
  <div>
    <v-card
        class="mx-auto"
    >
    <div class="text-center">
      <h2 class="purple--text">Packages</h2>
    </div>
      <v-container>
        <v-row dense>
          <v-col
            v-for="(item, i) in storedInfo"
            :key="i"
            cols="2"
            >
            <v-card
                :color="item.color"
                dark
            >
              <div class="d-flex flex-no-wrap justify-space-between">
                <div>
                  <v-card-title
                    class="text-h5"
                    v-text="item.name"/>
                  <v-card-subtitle 
                    v-text="'Cost '+item.cost+'$'"/>
                  <v-card-subtitle 
                    v-for="item2 in item.items" :key="item2" 
                    v-text="item2"/>
                </div>
              </div>
              <v-avatar>
                <v-img :src="item.img"/>
              </v-avatar>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-card>
  </div>
</template>

<script>
import axios from "axios";
import { mapState } from "vuex";
  export default {
    name:"Home",
    computed:{
      ...mapState(["host","storedInfo"])
    },
    created() {
        axios
        .get(this.host + "/list/planInfo")
        .then((response) => {
          this.storedInfo=response.data.data
          this.$store.commit("setStoredInfo", response.data.data);
          for(let prop in this.storedInfo){
            this.storedInfo[prop].img=require('../assets/'+this.storedInfo[prop].unique+'.png')
          }
        });
    },
  }
</script>