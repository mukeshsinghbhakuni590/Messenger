

var app = angular.module('myApp', ['ui.bootstrap']);
app.controller('myCtrl', function($scope, $http, $timeout, $interval, $q) {

  $scope.game_service_url =  localStorage.getItem("serverAddress")             // "http://127.0.0.1/gameservice"

  $scope.login = function() {
    $http(
         {
             url: $scope.game_service_url + '/userValidation',
             method: 'POST',
             headers : {
               'Content-Type':'application/json'
             },
             data : {
                 Email : $scope.email,
                 Passwd : $scope.pwd
             }
         }
    ).then(
        function (response) {
            console.log(response.data)
            localStorage.setItem("token", response['data']['_id'])
            localStorage.setItem("usrid", response['data']['usrid'])
            window.open("user_profile.html","_self")
        }
    )
    }
});
