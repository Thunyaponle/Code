
<?php
require 'productConnection.php';
if(isset($_POST["confirm"])){
    $sql = "SELECT * FROM Orders";
    $result = mysqli_query($conn, $sql);

    while ($row = mysqli_fetch_assoc($result)) {
    $pay = $row['status_pay'];
    if($pay=='Yes') {
    $ID = $row['order_id'];
    echo "$ID";
    $query2 = "UPDATE Orders SET status_served = 'Yes'where order_id = $ID";
    mysqli_query($conn,$query2);
    }
    }
  echo
  "
  <script> alert('Data Update Successfully'); </script>
  ";
} 
?>
<!DOCTYPE html>
<html>
<meta name="viewport" content="width=device-width, initial-scale=1">
<!--link rel="stylesheet" href="https://www.w3schools.com/w3css/4/w3.css"-->
<head>
<title>Admin order</title>
<style media="screen">
    label {
        display: block;
    }

    .mp-container {
        position: -webkit-sticky;
        position: sticky;
        top: 0;
        padding: 5px;
        background-color: #630515;
        border: 2px solid #b84e7c;
        text-align: center
    }

    .mp-sidebar {
        height: 100%;
        width: 200px;
        background-color: #df3434;
        position: fixed !important;
        z-index: 1;
        overflow: auto
    }

    .mp-bar-item {
        width: 100%;
        display: block;
        padding: 8px 16px;
        text-align: left;
        border: none;
        white-space: normal;
        float: none;
        outline: 0;
        color: #824848;
    }

    .mp-light,
    .mp-hover-light:hover {
        background-color: #f3d1d1 !important
    }

    .mp-block {
        display: block;
        width: 100%
    }

    .mp-bar-block {
        text-align: center
    }

    .w3-block {
        display: block;
        width: 100%
    }

    .mp-button:hover {
        color: #e1db58 !important;
        background-color: #c9768c !important
    }

    table, th, td {
        background-color: rgb(232, 234, 192);
        border: 5px solid black;
        border-collapse: collapse;
        text-align: center;
        
        } 
        
        th,td{ padding: 15px;}

    .pad {padding: 15px;}
       
</style>
</head>

<body>

    <!-- Sidebar -->
    <div class="mp-sidebar mp-light  mp-bar-block "  style="width:25%">
        <h3 class="mp-bar-item">เมนู</h3>
        <a href="admin_product.php" class="mp-bar-item mp-button">คลังสินค้า</a>
        <a href="admin_order.php" class="mp-bar-item mp-button">รายการอาหาร</a>
        <a href="admin_served.php" class="mp-bar-item mp-button">เสิร์ฟ</a>
       
    </div>

    <!-- Page Content -->
    <div style="margin-left:25%">

        <div class="mp-container">
            <h1 style="color:rgb(255, 255, 255)"="">รายการอาหารที่จ่ายแล้ว</h1>
        </div>

        <div class="w3-container">
            <center>

        <br>

<br>
<br>
<center><h2>Bills</h2></center>      
<?php
require 'productConnection.php';
$sql="select o.order_id,r.riceName,m.main_courseName,si.side_dishesName,sn.snacksName,d.dessertName,o.total_price,o.status_pay,o.status_served from Orders o 
left join rice r on o.riceID = r.riceID
left join main_course m on o.main_courseID = m.main_courseID
left join side_dishes si on o.side_dishesID = si.side_dishesID
left join snacks sn on o.snacksID = sn.snacksID
left join dessert d on o.dessertID = d.dessertID
where o.status_served = 'Cooking';";

$result = $conn->query($sql);
if($result->num_rows > 0){
   echo "<table border=1>";
   echo "<tr><td>No.</td><td>Rice</td><td>Main_course</td><td>Side_dishes</td><td>Snacks</td><td>Dessert</td><td>Total price</td><td>Payment status</td><td>Serving status</td></tr>";


   while($row=$result->fetch_assoc()){
     echo "<tr><td>".$row['order_id']."</td><td>".$row['riceName']."</td><td>".$row['main_courseName']."</td><td>".$row['side_dishesName']."</td><td>".$row['snacksName']."</td><td>".$row['dessertName']."</td><td>".$row['total_price']."</td><td>".$row['status_pay']."</td><td>".$row['status_served']."</td></tr>";
   }
      
   echo "</table>";
}else{
   echo "0 rows available";
}
$conn->close();;
?>
<br>
<br>
<center><h2>Update</h2></center>
<?php
require 'productConnection.php';
$sql="select o.order_id,r.riceName,m.main_courseName,si.side_dishesName,sn.snacksName,d.dessertName,o.total_price,o.status_pay,o.status_served from Orders o 
left join rice r on o.riceID = r.riceID
left join main_course m on o.main_courseID = m.main_courseID
left join side_dishes si on o.side_dishesID = si.side_dishesID
left join snacks sn on o.snacksID = sn.snacksID
left join dessert d on o.dessertID = d.dessertID
where o.status_served = 'Yes';";

$result = $conn->query($sql);
if($result->num_rows > 0){
   echo "<table border=1>";
   echo "<tr><td>No.</td><td>Rice</td><td>Main_course</td><td>Side_dishes</td><td>Snacks</td><td>Dessert</td><td>Total price</td><td>Payment status</td><td>Serving status</td></tr>";


   while($row=$result->fetch_assoc()){
     echo "<tr><td>".$row['order_id']."</td><td>".$row['riceName']."</td><td>".$row['main_courseName']."</td><td>".$row['side_dishesName']."</td><td>".$row['snacksName']."</td><td>".$row['dessertName']."</td><td>".$row['total_price']."</td><td>".$row['status_pay']."</td><td>".$row['status_served']."</td></tr>";
   }
      
   echo "</table>";
}else{
   echo "0 rows available";
}
$conn->close();;
?>
                <br>
                <form class="" action="" method="post" autocomplete="off">
                <button type="submit" name="confirm">Confirm served</button>
                </form>
            </center>
           
        </div>
    

    </div>


    
   

</body>

</html>