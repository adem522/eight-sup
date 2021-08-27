<template>
  <v-container
    fluid
    style="
      width: 100%;
      display: flex;
      align-items: center;
      justify-content: flex-start;
      flex-direction: column;
      margin: 0 auto;
    "
  >
    <v-card
        class="mx-auto"
        max-width="800"
        tile
        v-for="item in event" :key="item.date"
    >
        <v-list-item three-line >
            <v-list-item-content>
                <v-list-item-title>Package {{item.unique}}</v-list-item-title>
                <v-list-item-subtitle v-if="type=='streamer'">Buyer Username {{item.buyerUsername}}</v-list-item-subtitle>
                <v-list-item-subtitle v-if="type=='client'">Seller Username {{item.sellerUsername}}</v-list-item-subtitle>
                <v-list-item-subtitle>Date {{item.date}}</v-list-item-subtitle>
            </v-list-item-content>
        </v-list-item>
    </v-card>
  </v-container>
</template>

<script>
import axios from 'axios';
import { mapState } from "vuex";
export default {
    name: "History",
    data(){
        return{
            event:[],
        };        
    },
    computed:{
        ...mapState(["username","host","storedInfo","type"])
    },
    created() {
        /*if(this.getEvent.length==0)*/ this.getEvent()
    },
    methods:{
        getEvent(){
            axios
            .put(this.host+"/list/event",{
                username:this.username,
                type:this.type,
            })
            .then((response) => {
                this.event=response.data
            })
        }
    }
}
</script>