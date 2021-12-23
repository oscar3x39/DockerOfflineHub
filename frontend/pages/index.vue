<template>
  <v-row justify="center" align="center">
    <v-col cols="12" sm="8" md="12">
      <v-autocomplete
        v-model="model"
        color="white"
        item-text="Description"
        item-value="API"
        label="Type dockerhub image name"
        placeholder="ex: mysql:5.7"
        prepend-icon="mdi-database-search"
        return-object
      ></v-autocomplete>
      <v-btn
        elevation="2"
      >Download Image</v-btn>
    </v-col>
    <v-col cols="12" sm="8" md="12">
      <iframe src=""></iframe>
    </v-col>
  </v-row>
</template>

<script>
export default {
  name: 'IndexPage'
}
</script>
