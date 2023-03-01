<template>
    <div id="content">
        <div v-if="field != null" id="game">
            <div v-for="(i, x) in field" :key="x" class="row">
                <div
                    class="cell"
                    :class="colorCell(cell)"
                    v-for="(cell, y) in i" 
                    :key="y" 
                    @click="leftClick(x,y)" 
                    @contextmenu.prevent="rightClick(x,y)"
                    >
                    <div v-if="!cell.flagged">{{ decodeCell(cell) }}</div>
                    <v-icon v-else>mdi-flag</v-icon>
                </div>
            </div>
        </div>
        <ChatCard id="chat" ref="chat" @send="chat"></ChatCard>
    </div>
</template>

<script>
    import ChatCard from "@/components/chatCard.vue"
    import {fetchField} from "@/http/http"
    export default {
        name: 'MineField',
        components:{ChatCard},

        data: () => ({
            field: null,
            socket: WebSocket.prototype,
        }),
        methods: {
            rightClick: function(_x,_y){
                this.send("rightClick", {x:_x, y:_y})
            },
            leftClick: function(_x,_y){
                this.send("leftClick", {x:_x, y:_y})
            },
            send: function(type, data){
                let msg = {
                    msgType : type, // Message type
                    msgData : data, // Data server will use
                }

                // Send it
                this.socket.send(JSON.stringify(msg))
            },
            chat: function(msg){
                this.send("chat", msg)
            },
            receiveChat(msg){
                this.$refs.chat.receive(msg)
            },
            decodeCell(cell){
                if(cell.revealed){
                    return cell.count == 0 ? "": cell.count
                }
            },
            colorCell(cell){
                if (!cell.revealed)
                    return cell.flagged? "flagged" :"covered"

                return "c"+cell.count
            },
            handleMessage(data){
                const msg = JSON.parse(data)

                switch(msg.msgType){
                    case "Update": this.field = msg.msgData; break;
                    case "chat": this.receiveChat(msg.msgData); break;
                }
            },
            JoinLobby(){
                const host = process.env.NODE_ENV === 'development' ? "localhost:8080" : window.location.host;
                const url = `ws://${host}/api/join`

                this.socket = new WebSocket(url)

                this.socket.onopen = function(){
                    this.connected = true;
                    console.log("Connected!")
                }.bind(this)

                this.socket.onclose = function() {
                    console.log("Disconnected!")
                }

                this.socket.onmessage = function(e){
                    this.handleMessage(e.data)
                }.bind(this);
            },

        },
        created: async function() {
            this.JoinLobby()
            this.field = await fetchField()
        }
    };
</script>


<style>

#content{
    display: flex;
    flex-wrap: nowrap;
    justify-content: space-between;
}

#chat{
    position: fixed;
    margin-top: 40px;
    left: 20px;
}

#game{
    border: red 2p solid;
    margin: 50px;
    margin-left: auto;
    margin-right: auto;
    padding: 1px;
    border-radius: 10px;
    width: 50%;
}
.cell{
    width: 38px;
    height: 38px;
    text-align: center;
    vertical-align: middle;
    line-height: 38px;       /* The same as your div height */
    border-radius: 3px;
}

.covered{
    background-color: #4C00B0;
}

.flagged{
    background-color: darkred;
}

.covered:hover, .flagged:hover{
    cursor: pointer;
}


/* Not great but whatever */
.c0{
    background-color: #121212;
}

.c1{
    background-color: #913175;
}

.c2{
    background-color: #CD5888;
}

.c3{
    background-color: #42d6a4;
}

.c4{
    background-color: #DC5F00;
}

.c5{
    background-color: #060047;
}

.c6{
    background-color: #B3005E;
}

.c7{
    background-color: #F2CD5C;
}

.c8{
    background-color: #46C2CB;
}

.row{
    display: flex;
    flex-wrap: nowrap;
}
</style>