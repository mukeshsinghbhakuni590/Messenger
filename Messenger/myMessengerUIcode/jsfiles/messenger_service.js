
var app = angular.module('myApp', ['ui.bootstrap']);
app.controller('myCtrl', function($scope, $http, $timeout, $interval, $q) {

  $scope.messenger_service_url = "http://127.0.0.1/myMessenger" //"http://192.168.43.81/gameservice"
  localStorage.setItem("serverAddress", $scope.messenger_service_url )

  $scope.create_user = function() {
         $http(
              {
                  url: $scope.messenger_service_url + '/user',
                  method: 'POST',
                  headers : {
                    'Content-Type':'application/json'
                  },
                  data : {
                      Uname : $scope.name,
                      Email : $scope.email,
                      Passwd : $scope.pwd
                  }
              }
      ).then(
              function (response) {
                  console.log(response)
                  window.open("htmlfiles/user_validation.html","_self")
              }
      )
  }

  $scope.login_link = function () {
    window.open("htmlfiles/user_validation.html","_self");
    }
});
