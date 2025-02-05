<?php
require 'productConnection.php';
if (isset($_POST["onikiri"])) {

    $query = "UPDATE rice SET countrice = 10 WHERE riceID=101";

    mysqli_query($conn, $query);
    echo
    "
  <script> alert('Data Inserted Successfully'); </script>
  ";
} else if (isset($_POST["seaweed_rice"])) {
    $query = "UPDATE rice SET countrice = 10 WHERE riceID=102";
    mysqli_query($conn, $query);
    echo
    "
  <script> alert('Data Inserted Successfully'); </script>
  ";
} else if (isset($_POST["japan_rice"])) {
    $query = "UPDATE rice SET countrice = 10 WHERE riceID=103";
    mysqli_query($conn, $query);
    echo
    "
  <script> alert('Data Inserted Successfully'); </script>
  ";
} else if (isset($_POST["fried_rice"])) {
    $query = "UPDATE rice SET countrice = 10 WHERE riceID=104";
    mysqli_query($conn, $query);
    echo
    "
  <script> alert('Data Inserted Successfully'); </script>
  ";
} else if (isset($_POST["tonkatsu"])) {
    $query = "UPDATE main_course SET countmain_course = 10 WHERE main_courseID=201";
    mysqli_query($conn, $query);
    echo
    "
  <script> alert('Data Inserted Successfully'); </script>
  ";
} else if (isset($_POST["teriyaki"])) {
    $query = "UPDATE main_course SET countmain_course = 10 WHERE main_courseID=202";
    mysqli_query($conn, $query);
    echo
    "
  <script> alert('Data Inserted Successfully'); </script>
  ";
} else if (isset($_POST["chicken_tonkatsu"])) {
    $query = "UPDATE main_course SET countmain_course = 10 WHERE main_courseID=203";
    mysqli_query($conn, $query);
    echo
    "
  <script> alert('Data Inserted Successfully'); </script>
  ";
} else if (isset($_POST["tempura"])) {
    $query = "UPDATE main_course SET countmain_course = 10 WHERE main_courseID=204";
    mysqli_query($conn, $query);
    echo
    "
  <script> alert('Data Inserted Successfully'); </script>
  ";
} else if (isset($_POST["seaweed_salad"])) {
    $query = "UPDATE side_dishes SET countside_dishes = 10 WHERE side_dishesID=301";
    mysqli_query($conn, $query);
    echo
    "
  <script> alert('Data Inserted Successfully'); </script>
  ";
} else if (isset($_POST["kimchi"])) {
    $query = "UPDATE side_dishes SET countside_dishes = 10 WHERE side_dishesID=302";
    mysqli_query($conn, $query);
    echo
    "
  <script> alert('Data Inserted Successfully'); </script>
  ";
} else if (isset($_POST["pickled_Radish"])) {
    $query = "UPDATE side_dishes SET countside_dishes = 10 WHERE 	side_dishesID=303";
    mysqli_query($conn, $query);
    echo
    "
  <script> alert('Data Inserted Successfully'); </script>
  ";
} else if (isset($_POST["Spinach_Sesame_Oil"])) {
    $query = "UPDATE side_dishes SET countside_dishes = 10 WHERE 	side_dishesID=304";
    mysqli_query($conn, $query);
    echo
    "
  <script> alert('Data Inserted Successfully'); </script>
  ";
} else if (isset($_POST["sausage"])) {
    $query = "UPDATE snacks SET countsnacks = 10 WHERE snacksID=401";
    mysqli_query($conn, $query);
    echo
    "
  <script> alert('Data Inserted Successfully'); </script>
  ";
} else if (isset($_POST["egg_rolls"])) {
    $query = "UPDATE snacks SET countsnacks = 10 WHERE snacksID=402";
    mysqli_query($conn, $query);
    echo
    "
  <script> alert('Data Inserted Successfully'); </script>
  ";
} else if (isset($_POST["karaage"])) {
    $query = "UPDATE snacks SET countsnacks = 10 WHERE snacksID=403";
    mysqli_query($conn, $query);
    echo
    "
  <script> alert('Data Inserted Successfully'); </script>
  ";
} else if (isset($_POST["takoyaki"])) {
    $query = "UPDATE snacks SET countsnacks = 10 WHERE snacksID=404";
    mysqli_query($conn, $query);
    echo
    "
  <script> alert('Data Inserted Successfully'); </script>
  ";
} else if (isset($_POST["dango"])) {
    $query = "UPDATE dessert SET countdessert = 10 WHERE dessertID=501";
    mysqli_query($conn, $query);
    echo
    "
  <script> alert('Data Inserted Successfully'); </script>
  ";
} else if (isset($_POST["mochi"])) {
    $query = "UPDATE dessert SET countdessert = 10 WHERE dessertID=502";
    mysqli_query($conn, $query);
    echo
    "
  <script> alert('Data Inserted Successfully'); </script>
  ";
} else if (isset($_POST["pudding"])) {
    $query = "UPDATE dessert SET countdessert = 10 WHERE dessertID=503";
    mysqli_query($conn, $query);
    echo
    "
  <script> alert('Data Inserted Successfully'); </script>
  ";
} else if (isset($_POST["castella"])) {
    $query = "UPDATE dessert SET countdessert = 10 WHERE dessertID=504";
    mysqli_query($conn, $query);
    echo
    "
  <script> alert('Data Inserted Successfully'); </script>
  ";
}
?>
<!DOCTYPE html>
<html>
<meta name="viewport" content="width=device-width, initial-scale=1">
<!--link rel="stylesheet" href="https://www.w3schools.com/w3css/4/w3.css"-->

<head>
    <title>Admin page</title>
    <style media="screen">
        label {
            display: block;
        }

        .mp-container {
            position: -webkit-sticky;
            position: sticky;
            top: 0;
            padding: 1px;
            background-color: #630111;
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
            color: #e1db18 !important;
            background-color: #c9768c !important
        }

        table,
        th,
        td {
            background-color: rgb(232, 234, 192);
            border: 1px solid black;
            border-collapse: collapse;
            text-align: center;

        }

        th,
        td {
            padding: 11px;
        }

        .pad {
            padding: 11px;
        }
    </style>
</head>

<body>

    <!-- Sidebar -->
    <div class="mp-sidebar mp-light  mp-bar-block " style="width:21%">
        <h3 class="mp-bar-item">เมนู</h3>
        <a href="admin_product.php" class="mp-bar-item mp-button">คลังสินค้า</a>
        <a href="admin_order.php" class="mp-bar-item mp-button">รายการอาหาร</a>
        <a href="admin_served.php" class="mp-bar-item mp-button">เสิร์ฟ</a>

    </div>

    <!-- Page Content -->
    <div style="margin-left:21%">

        <div class="mp-container">
            <h1 style="color:rgb(255, 255, 255)"="">รายการวัตถุดิบ</h1>
        </div>

        <div class="w3-container">
            <center>

                <br>
                <center>
                    <h2>ข้าว</h2>
                </center>
                <?php
                require 'productConnection.php';
                $sql1 = "SELECT * FROM rice;";
                $result1 = $conn->query($sql1);

                if ($result1->num_rows > 0) {
                    echo "<table border=1>";
                    echo "<tr><td>No.</td><td>Name</td><td>Total</td><td>Price</td></tr>";
                    while ($row = $result1->fetch_assoc()) {
                        echo "<tr><td>" . $row['riceID'] . "</td><td>" . $row['riceName'] . "</td><td>" . $row['countrice'] . "</td><td>" . $row['rice_price'] . "</td></tr>";
                    }
                    echo "</table>";
                } else {
                    echo "0 rows available";
                }
                $conn->close();
                ?>
                <br>
                <center>
                    <h2>อาหารจานหลัก</h2>
                </center>
                <?php
                require 'productConnection.php';
                $sql2 = "SELECT * FROM main_course;";
                $result2 = $conn->query($sql2);

                if ($result2->num_rows > 0) {
                    echo "<table border=1>";
                    echo "<tr><td>No.</td><td>Name</td><td>Total</td><td>Price</td></tr>";
                    while ($row = $result2->fetch_assoc()) {
                        echo "<tr><td>" . $row['main_courseID'] . "</td><td>" . $row['main_courseName'] . "</td><td>" . $row['countmain_course'] . "</td><td>" . $row['main_course_price'] . "</td></tr>";
                    }
                    echo "</table>";
                } else {
                    echo "0 rows available";
                }
                $conn->close();
                ?>
                <br>
                <center>
                    <h2>เครื่องเคียง</h2>
                </center>
                <?php
                require 'productConnection.php';
                $sql2 = "SELECT * FROM side_dishes;";
                $result2 = $conn->query($sql2);

                if ($result2->num_rows > 0) {
                    echo "<table border=1>";
                    echo "<tr><td>No.</td><td>Name</td><td>Total</td><td>Price</td></tr>";
                    while ($row = $result2->fetch_assoc()) {
                        echo "<tr><td>" . $row['side_dishesID'] . "</td><td>" . $row['side_dishesName'] . "</td><td>" . $row['countside_dishes'] . "</td><td>" . $row['side_dishes_price'] . "</td></tr>";
                    }
                    echo "</table>";
                } else {
                    echo "0 rows available";
                }
                $conn->close();
                ?>
                <br>
                <center>
                    <h2>ของทานเล่น</h2>
                </center>
                <?php
                require 'productConnection.php';
                $sql4 = "SELECT * FROM snacks;";
                $result4 = $conn->query($sql4);

                if ($result4->num_rows > 0) {
                    echo "<table border=1>";
                    echo "<tr><td>No.</td><td>Name</td><td>Total</td><td>Price</td></tr>";
                    while ($row = $result4->fetch_assoc()) {
                        echo "<tr><td>" . $row['snacksID'] . "</td><td>" . $row['snacksName'] . "</td><td>" . $row['countsnacks'] . "</td><td>" . $row['snacks_price'] . "</td></tr>";
                    }
                    echo "</table>";
                } else {
                    echo "0 rows available";
                }
                $conn->close();
                ?>
                <br>
                <center>
                    <h2>ของหวาน</h2>
                </center>
                <?php
                require 'productConnection.php';
                $sql1 = "SELECT * FROM dessert;";
                $result1 = $conn->query($sql1);

                if ($result1->num_rows > 0) {
                    echo "<table border=1>";
                    echo "<tr><td>No.</td><td>Name</td><td>Total</td><td>Price</td></tr>";
                    while ($row = $result1->fetch_assoc()) {
                        echo "<tr><td>" . $row['dessertID'] . "</td><td>" . $row['dessertName'] . "</td><td>" . $row['countdessert'] . "</td><td>" . $row['dessert_price'] . "</td></tr>";
                    }
                    echo "</table>";
                } else {
                    echo "0 rows available";
                }
                $conn->close();
                ?>

                <br><br>
                <center>
                    <h2>ตารางเพิ่มวัตถุดิบ</h2>
                </center>
                <form class="" action="" method="post" autocomplete="off">
                    <table>
                        <tr>
                            <td>
                                ข้าว
                            </td>
                            <td>
                                <button type="submit" name="onikiri">add Onikiri</button>
                            </td>
                            <td>
                                <button type="submit" name="seaweed_rice">add Seaweed rice</button>
                            </td>
                            <td>
                                <button type="submit" name="japan_rice">add Japan rice</button>
                            </td>
                            <td>
                                <button type="submit" name="fried_rice">add fried rice</button>
                            </td>
                        </tr>
                        <tr>
                            <td>
                                อาหารจานหลัก
                            </td>
                            <td>
                                <button type="submit" name="tonkatsu">add Tonkatsu</button>
                            </td>
                            <td>
                                <button type="submit" name="teriyaki">add Teriyaki</button>
                            </td>
                            <td>
                                <button type="submit" name="chicken_tonkatsu">add Chicken tonkatsu</button>
                            </td>
                            <td>
                                <button type="submit" name="tempura">add Tempura</button>
                            </td>
                        </tr>
                        <tr>
                            <td>
                                เครื่องเคียง
                            </td>
                            <td>
                                <button type="submit" name="seaweed_salad">add Seaweed Salad</button>
                            </td>
                            <td>
                                <button type="submit" name="kimchi">add Kimchi</button>
                            </td>
                            <td>
                                <button type="submit" name="pickled_Radish">add Pickled Radish</button>
                            </td>
                            <td>
                                <button type="submit" name="Spinach_Sesame_Oil">add Spinach Sesame Oil</button>
                            </td>
                        </tr>
                        <td>
                            ของทานเล่น
                        </td>
                        <td>
                            <button type="submit" name="sausage">add Sausage</button>
                        </td>
                        <td>
                            <button type="submit" name="egg_rolls">add Egg rolls</button>
                        </td>
                        <td>
                            <button type="submit" name="karaage">add Karaage</button>
                        </td>
                        <td>
                            <button type="submit" name="takoyaki">add Takoyaki</button>
                        </td>
                        </tr>
                        <td>
                            ของหวาน
                        </td>
                        <td>
                            <button type="submit" name="dango">add Dango</button>
                        </td>
                        <td>
                            <button type="submit" name="mochi">add Mochi</button>
                        </td>
                        <td>
                            <button type="submit" name="pudding">add Pudding</button>
                        </td>
                        <td>
                            <button type="submit" name="castella">add Castella</button>
                        </td>
                        </tr>
                    </table>
                </form>
            </center>
            <br>
            <br>
        </div>


    </div>





</body>

</html>
admin_product.php
17 KB