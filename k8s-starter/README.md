# myshop

## Helm
```
$ helm create myshop
Creating myshop
```


## Connect from Pod to Pod

```
"mongodb://mongodb-cluster-ip-service/products",
```

mongodb-cluster-ip-service is service name.

## Ingress Controller
[For docker desktop](https://kubernetes.github.io/ingress-nginx/deploy/#quick-start)
[For GKE](https://kubernetes.github.io/ingress-nginx/deploy/#gce-gke)

## TestData
```
// http://localhost/api/products POST

{
    "name": "Red Velvet",
    "image": "https://media.istockphoto.com/photos/red-velvet-cake-picture-id485832764?k=20&m=485832764&s=612x612&w=0&h=58yxSE0lOx3aD7OZJ3wodakdAbEB-aH6MWNj3QGutwg=",
    "description": "Moist with cream cheese icing",
    "price": 34.33
}

{
    "name": "Raspberry Cheesecake",
    "image": "https://www.elmundoeats.com/wp-content/uploads/2020/06/FP-No-Bake-Raspberry-Cheesecake.jpg",
    "description": "New York style cheesecake",
    "price": 64.33
}
```
