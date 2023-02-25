<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    >
    <h1>
      MINES
    </h1>
    </v-app-bar>

    <v-main>
      <MineField/>
    </v-main>
  </v-app>
</template>

<script>
import MineField from './components/MineField';

export default {
  name: 'App',

  components: {
    MineField,
  },

  data: () => ({
    //
  }),
};
</script>
