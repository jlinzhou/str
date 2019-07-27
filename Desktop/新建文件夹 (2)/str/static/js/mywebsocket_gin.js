/*
author:Zhou Junlin
time:2019/06/13
company:shangyi
 */
$(function () {
    url = 'ws://localhost:3000/ping';
    c = new WebSocket(url);

    send = function(data){
    //$("#output").append((new Date())+ " ==> "+data+"\n")
   // console.log((new Date())+ " ==> "+data);
        c.send(data)
    };

    c.onmessage = function(msg){
    //$("#output").append((new Date())+ " <== "+msg.data+"\n")
       // console.log(msg)
    };

    c.onopen = function(){
        setInterval(
        function(){ send("ping") }
        , 1000 )
    };
     function init(m_onorder)
    {
        var onorder= m_onorder;//{{ .onorder }};
        console.log( onorder.Value );
        //$('.onrtnorder')

        for(var i=0;i<onorder.Value.length;i++)
        {
            var orderdata=onorder.Value[i];
            orderdata='<tr>\
                <td>'+onorder.Value[i]['Investorid']+'</td>\
                <td>'+orderdata['Userid']+'</td>\
                <td>'+orderdata['Ordersysid']+'</td>\
                <td>'+orderdata['Instrumentid']+'</td>\
                <td>'+orderdata['Direction']+'</td>\
                <td>'+orderdata['Offsetflag']+'</td>\
                <td>'+parseFloat(orderdata['Limitprice'])+'</td>\
                <td>'+orderdata['Volume']+'</td>\
                <td>'+orderdata['Orderstatus']+'</td>\
                <td>'+orderdata['Volumetraded']+'</td>\
                <td>'+orderdata['Volumecancled']+'</td>\
                <td>'+orderdata['Volumeremain']+'</td>\
                <td>'+orderdata['Inserttime']+'</td>\
            </tr>';

            $('.onrtnorder').children().eq(1).append(orderdata);
        }
    }
});