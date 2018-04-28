
var app = angular.module('myApp', ['ui.bootstrap']);
app.controller('myCtrl', function($scope, $http, $timeout, $interval, $q) {

    $scope.messenger_service_url =localStorage.getItem("serverAddress") //"http://127.0.0.1/gameservice"

    $scope.usrid = localStorage.getItem("usrid")
    $scope.usrtkn = localStorage.getItem("token")
    $scope.userinfo = []
    $scope.sviews = []
    $scope.gviews = []
    $scope.contid = null
    $scope.continfo = []

    $scope.currviewid = null
    $scope.currviewinfo = []
    $scope.msg = []
    $scope.othercontinfo = []
    $scope.currgrpid = null
    $scope.currgrpinfo = []
    //--------------------------------------------------------
    $scope.checkmemberlist = []

    $scope.logout = function(){
        $http(
            {
                url: $scope.messenger_service_url + '/userValidation',
                method: 'PUT',
                headers : {
                    'Content-Type':'application/json'
                },
                data : {
                    Session_id : $scope.usrtkn,
                    Email : userinfo.email,
                    User_id : $scope.usrid
                }
            }
        ).then(
            function (response) {
                console.log(response.data)
                localStorage.removeItem("usrid")
                localStorage.removeItem("token")
                window.open("../messenger_service.html","_self")
            }
        )
    }



    $scope.get_user_info = function() {
        $http(
            {
                url: $scope.messenger_service_url + '/user',
                method: 'GET',
                params: {
                    sessionid: $scope.usrtkn,
                    userid : $scope.usrid
                }
            }
        ).then(
            function (response) {
                console.log(response)
                $scope.userinfo = response.data
                $scope.contid = $scope.userinfo.cont_id
                $scope.get_cont_info()
            }
        )
    }

    $scope.get_cont_info = function() {
        $http(
            {
                url: $scope.messenger_service_url + '/user',
                method: 'PUT',
                params: {
                    operation : "continfo"
                },
                headers : {
                    'Content-Type':'application/json'
                },
                data : {
                    Session_id : $scope.usrtkn,
                    cont_id : $scope.contid
                }
            }
        ).then(
            function (response) {
                console.log(response.data)
                $scope.continfo = response.data
                $scope.sviews = $scope.continfo.sviews
                $scope.gviews = $scope.continfo.gviews
            }
        )
    }

    $scope.change_user_name = function() {
        $http(
            {
                url: $scope.messenger_service_url + '/user',
                method: 'PUT',
                params: {
                    operation : "changeName"
                },
                headers : {
                    'Content-Type':'application/json'
                },
                data : {
                    User_id      : $scope.usrid,
                    New_User_name : $scope.newusername,
                    Session_id : $scope.usrtkn,
                }
            }
        ).then(
            function (response) {
                console.log(response.data)
                $scope.userinfo = response.data
            }
        )
    }
    $scope.change_user_status = function() {
        $http(
            {
                url: $scope.messenger_service_url + '/user',
                method: 'PUT',
                params: {
                    operation : "changeStatus"
                },
                headers : {
                    'Content-Type':'application/json'
                },
                data : {
                    User_id      : $scope.usrid,
                    New_User_status : $scope.newuserstatus,
                    Session_id : $scope.usrtkn,
                }
            }
        ).then(
            function (response) {
                console.log(response.data)
                $scope.userinfo = response.data
            }
        )
    }
    $scope.change_user_passwd = function() {
        $http(
            {
                url: $scope.messenger_service_url + '/user',
                method: 'PUT',
                params: {
                    operation : "changePasswd"
                },
                headers : {
                    'Content-Type':'application/json'
                },
                data : {
                    User_id      : $scope.usrid,
                    Old_passwd  : $scope.useroldpasswd,
                    New_passwd  : $scope.usernewpasswd,
                    Session_id : $scope.usrtkn,
                }
            }
        ).then(
            function (response) {
                console.log(response.data)
                $scope.userinfo = response.data
            }
        )
    }
    $scope.add_contact = function() {
        $http(
        {
            url: $scope.messenger_service_url + '/user',
            method: 'PUT',
            params: {
                operation : "addCont"
            },
            headers : {
                'Content-Type':'application/json'
            },
            data : {
                Self_cont_id : $scope.contid,
                Member_cont_id : $scope.membercontid,
                Session_id : $scope.usrtkn,
            }
        }
        ).then(
            function (response) {
                console.log(response.data)
                $scope.continfo = response.data
                $scope.sviews = $scope.continfo.sviews
                $scope.gviews = $scope.continfo.gviews
            }
        )
    }
    $scope.block_contact = function() {
        $http(
            {
                url: $scope.messenger_service_url + '/user',
                method: 'PUT',
                params: {
                    operation : "blockCont"
                },
                headers : {
                    'Content-Type':'application/json'
                },
                data :    {
                    Self_cont_id : $scope.contid,
                    Member_cont_id : $scope.membercontid,
                    Session_id : $scope.usrtkn,
                }
            }
        ).then(
            function (response) {
                console.log(response.data)
                $scope.continfo = response.data
                $scope.sviews = $scope.continfo.sviews
                $scope.gviews = $scope.continfo.gviews
            }
        )
    }
    $scope.leave_group = function() {
        $http(
            {
                url: $scope.messenger_service_url + '/user',
                method: 'PUT',
                params: {
                    operation : "leaveGroup"
                },
                headers : {
                    'Content-Type':'application/json'
                },
                data :    {
                    Cont_id  :  $scope.contid,
                    Group_id :   $scope.grpid,
                    Session_id : $scope.usrtkn,
                }
            }
        ).then(
            function (response) {
                console.log(response.data)
                $scope.continfo = response.data
                $scope.sviews = $scope.continfo.sviews
                $scope.gviews = $scope.continfo.gviews
            }
        )
    }
    $scope.other_info = function() {
        $http(
            {
                url: $scope.messenger_service_url + '/user',
                method: 'PUT',
                params: {
                    operation : "otherinfo"
                },
                headers : {
                    'Content-Type':'application/json'
                },
                data : {
                    User_id : $scope.userid,
                    Other_cont_id : $scope.othercontid,
                    Session_id : $scope.usrtkn,
                }
            }
        ).then(
            function (response) {
                console.log(response.data)
                $scope.othercontinfo = response.data
            }
        )
    }
    $scope.show_view = function(cont_id,grp_id){
        var k = 0
        if (cont_id != "") {
            var len = $scope.sviews.length
            for (i=0; i<len; i++) {
                if ($scope.sviews[i].Cont_id == cont_id ) {
                    $scope.currviewid = $scope.sviews[i].View_id
                    $scope.checkmemberlist = []
                    $scope.checkmemberlist.push(cont_id)
                    $scope.get_view_by_id()
                    k = 1
                    break
                }
            }
            if (k==0) {
                $http(
                    {
                        url: $scope.messenger_service_url + '/view',
                        method : 'POST',
                        headers : {
                            'Content-Type':'application/json'
                        },
                        data : {
                            Msg_ids : [],
                            Session_id : $scope.usrtkn,
                            Self_cont_id : $scope.contid,
                            Other_cont_id : cont_id,
                        }
                    }
                ).then(
                    function (response) {
                        console.log(response.data)
                        $scope.currviewinfo = response.data
                        $scope.currviewid = $scope.currviewinfo._id
                        $scope.checkmemberlist = []
                        $scope.checkmemberlist.push(cont_id)
                        $scope.get_cont_info()
                    }
                )
            }
        } 
        else if (grp_id != "") {
            var len = $scope.gviews.length
            for (i=0; i<len; i++) {
                if ($scope.gviews[i].Cont_id == grp_id ) {
                    $scope.currviewid = $scope.gviews[i].View_id
                    $scope.checkmemberlist = []
                    $scope.get_grp_info_by_id(grp_id)
                    for (i=0;i<($scope.currgrpinfo.group_members.length);i++){
                        $scope.checkmemberlist.push($scope.currgrpinfo.group_members[i])
                    }
                    $scope.get_view_by_id()
                    break
                }
            }
        }
    }

    $scope.get_view_by_id = function(){
        $http(
           {
                url: $scope.messenger_service_url + '/view',
                method: 'GET',
                params: {
                    sessionid : $scope.usrtkn,
                    viewid : $scope.currviewid
                }
           }
        ).then(
            function (response) {
                console.log(response.data)
                $scope.currviewinfo = response.data
            }
        )
    }

    $scope.get_msg_by_id = function(msg_id){
        $http(
           {
                url: $scope.messenger_service_url + '/message',
                method: 'GET',
                params: {
                    sessionid : $scope.usrtkn,
                    msgid : msg_id
                }
           }
        ).then(
           function (response) {
                console.log(response.data)
                $scope.msg = response.data
           }
        )
    }

    $scope.create_msg = function() {
        var a = []
        a.push($scope.currviewid)
        $http(
            {
                url: $scope.messenger_service_url + '/message',
                method: 'POST',
                headers : {
                    'Content-Type':'application/json'
                },
                data :    {
                    Sender : $scope.contid,
                    Receivers : $scope.checkmemberlist,
                    Msg_type : "string",
                    Msg_value : $scope.msgarea,
                    Views :a,
                    Session_id : $scope.usrtkn
                }
            }
        ).then(
            function (response) {
                console.log(response.data)
                $scope.get_view_by_id()
                $scope.msgarea = null
            }
        )
    }

    $scope.delete_msg = function(msgid) {
        $http(
            {
                url: $scope.messenger_service_url + '/view',
                method: 'PUT',
                headers : {
                    'Content-Type':'application/json'
                },
                data :    {
                    View_id : $scope.currviewid,
                    Session_id : $scope.usrtkn,
                    Msg_id : msgid ,
                    Cont_id : $scope.contid
                }
            }
        ).then(
            function (response) {
                console.log(response.data)
                $scope.currviewinfo = response.data
            }
        )
    }

    $scope.get_grp_info_by_id = function(grp_id){
        $http(
           {
                url: $scope.messenger_service_url + '/group',
                method: 'GET',
                params: {
                    sessionid : $scope.usrtkn,
                    grpid : grp_id
                }
           }
        ).then(
           function (response) {
                console.log(response.data)
                $scope.currgrpid = grp_id
                $scope.currgrpinfo = response.data
           }
        )
    }

    $scope.make_admin = function(memid) {
        $http(
            {
                url: $scope.messenger_service_url + '/group',
                method: 'PUT',
                params: {
                    operation : "makeAdmin"
                },
                headers : {
                    'Content-Type':'application/json'
                },
                data :    {
                    Admin_id : $scope.contid,
                    Member_id : memid,
                    Group_id :   $scope.currgrpid,
                    Session_id : $scope.usrtkn,
                }
            }
        ).then(
            function (response) {
                console.log(response.data)
                $scope.currgrpinfo = response.data
            }
        )
    }
    
    $scope.add_member = function(memid) {
        $http(
            {
                url: $scope.messenger_service_url + '/group',
                method: 'PUT',
                params: {
                    operation : "addMember"
                },
                headers : {
                    'Content-Type':'application/json'
                },
                data :    {
                    Admin_id : $scope.contid,
                    Member_id : memid,
                    Group_id :   $scope.currgrpid,
                    Session_id : $scope.usrtkn,
                }
            }
        ).then(
            function (response) {
                console.log(response.data)
                $scope.currgrpinfo = response.data
            }
        )
    }

    $scope.rem_member = function(memid) {
        $http(
            {
                url: $scope.messenger_service_url + '/group',
                method: 'PUT',
                params: {
                    operation : "remMember"
                },
                headers : {
                    'Content-Type':'application/json'
                },
                data :    {
                    Admin_id : $scope.contid,
                    Member_id : memid,
                    Group_id :   $scope.currgrpid,
                    Session_id : $scope.usrtkn,
                }
            }
        ).then(
            function (response) {
                console.log(response.data)
                $scope.currgrpinfo = response.data
            }
        )
    }

    $scope.rename_group = function(newgroupname) {
        $http(
            {
                url: $scope.messenger_service_url + '/group',
                method: 'PUT',
                params: {
                    operation : "renameGroup"
                },
                headers : {
                    'Content-Type':'application/json'
                },
                data :    {
                    Member_id : $scope.contid,
                    New_group_name : newgroupname,
                    Group_id :   $scope.currgrpid,
                    Session_id : $scope.usrtkn,
                }
            }
        ).then(
            function (response) {
                console.log(response.data)
                $scope.currgrpinfo = response.data
            }
        )
    }

    $scope.create_group = function(groupname) {
        var a = []
        a.push($scope.contid)
        $http(
            {
                url: $scope.messenger_service_url + '/group',
                method: 'POST',
                headers : {
                    'Content-Type':'application/json'
                },
                data :    {
                    Gname : groupname,
                    Admins : a,
                    Gmembers : $scope.checkmemberlist,
                    Session_id : $scope.usrtkn,
                }
            }
        ).then(
            function (response) {
                console.log(response.data)
                $scope.currgrpinfo = response.data
                $scope.currgrpid = response.data._id
            }
        )
    }



    $scope.get_user_info()
});
