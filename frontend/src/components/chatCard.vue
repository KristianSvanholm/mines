<template>
    <v-card width="20%" height="75%">
        <v-container class="chatBox" ref="chat">
            <v-row :key="i" v-for="(msg, i) in messages">
                <label class="chat">{{msg.name}} - {{msg.message}}</label>
            </v-row>
        </v-container>
        <div>
            <v-text-field
                v-model="msg.message"
                label="Message"
                @keyup.enter="send"
            ></v-text-field>
        </div>
    </v-card>
</template>

<script>

    export default {
        name: 'ChatCard',

        data () {
            return {
                messages: [],
                msg: {
                    message: '',
                }
            }
        }, 
        methods: {
            // Receives all messages in lobby and adds them to the chat
            receive(msg){
                this.messages.push(msg)

                //Scroll to bottom
                this.$nextTick(() => {
                    this.$refs.chat.scrollTop = this.$refs.chat.scrollHeight
                })
            },

            // Sends client message to lobby
            send(){
                if (this.msg.message == "") return

                this.$emit('send', this.msg)
                this.msg.message = '';
            },
        }
    }

</script>

<style>
    .chatBox {
        height: 100%;
        overflow-y: scroll;
        overflow-x: hidden;
    }

    .chat {
        padding-left: 5px;
        text-align: left;
        width : 100%;
        text-overflow: ellipsis;
    }

</style>
