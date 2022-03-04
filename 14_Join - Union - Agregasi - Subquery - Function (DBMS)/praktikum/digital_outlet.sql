-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 03, 2022 at 05:18 PM
-- Server version: 10.4.22-MariaDB
-- PHP Version: 8.1.2

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `digital_outlet`
--

-- --------------------------------------------------------

--
-- Table structure for table `operators`
--

CREATE TABLE `operators` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `operators`
--

INSERT INTO `operators` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Telkomsel', '2022-03-03 14:42:46', NULL),
(2, 'Indosat', '2022-03-03 14:42:46', NULL),
(3, 'XL', '2022-03-03 14:42:46', NULL),
(4, 'AXIS', '2022-03-03 14:42:46', NULL),
(5, 'Smartfren', '2022-03-03 14:42:46', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `payment_methods`
--

CREATE TABLE `payment_methods` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `status` smallint(1) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `payment_methods`
--

INSERT INTO `payment_methods` (`id`, `name`, `status`, `created_at`, `updated_at`) VALUES
(1, 'Transfer Bank', 1, '2022-03-03 14:56:17', NULL),
(2, 'Kartu Kredit/Debit', 1, '2022-03-03 14:56:17', NULL),
(3, 'Dompet Digital', 1, '2022-03-03 14:56:17', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `product_type_id` int(11) NOT NULL,
  `operator_id` int(11) NOT NULL,
  `code` varchar(50) NOT NULL,
  `name` varchar(100) NOT NULL,
  `status` smallint(1) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `product_type_id`, `operator_id`, `code`, `name`, `status`, `created_at`, `updated_at`) VALUES
(3, 2, 1, 'TSELD15', 'OMG! Ketengan 2.5 GB sampai 3.7 GB', 1, '2022-03-03 14:53:17', NULL),
(4, 2, 1, 'TSELN20', 'Disney+ Hotstar! Disney+ Hotstar 3 GB MAXStream', 1, '2022-03-03 14:53:17', NULL),
(5, 2, 1, 'TSELB15', 'ZOOM 1 hari 500Mb', 1, '2022-03-03 14:53:17', NULL),
(6, 3, 4, 'AXVP50K', 'Voucher Pulsa AXIS Rp.50.000', 1, '2022-03-03 14:53:17', NULL),
(7, 3, 4, 'AXVGML50K', 'Voucher Game AXIS! 170 Diamond MLBB', 1, '2022-03-03 14:53:17', NULL),
(8, 3, 4, 'AXVGPM75K', 'Voucher Game AXIS! 325 UC PUBGM', 1, '2022-03-03 14:53:17', NULL),
(9, 1, 5, 'SFVGPM75K', 'Pulsa Smartfren Rp.75.000', 1, '2022-03-03 14:53:17', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `product_descriptions`
--

CREATE TABLE `product_descriptions` (
  `id` int(11) NOT NULL,
  `description` text DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `product_descriptions`
--

INSERT INTO `product_descriptions` (`id`, `description`, `created_at`, `updated_at`) VALUES
(3, 'Harga Rp.15.000, Kuota 2.5 GB hingga 3.7 GB OMG! KUOTA INTERNET SESUAI ZONA USER. Masa aktif 1 hari.', '2022-03-03 14:54:08', NULL),
(4, 'Harga Rp.20.000, 3GB MAXStream selama 1 bulan. Sudah termasuk langganan Disney+ Hotstar 1 bulan.', '2022-03-03 14:54:08', NULL),
(5, 'Harga Rp.15.000, Kuota Video Conference ZOOM 500Mb. Masa aktif 1 hari.', '2022-03-03 14:54:08', NULL),
(6, 'Harga Rp.53.000, Voucher Pulsa AXIS Rp.50.000. Berlaku 1 kali.', '2022-03-03 14:54:08', NULL),
(7, 'Harga Rp.50.000, Voucher Game AXIS! 170 Diamond Mobile Legends: Bang Bang. Berlaku untuk 1 akun.', '2022-03-03 14:54:08', NULL),
(8, 'Harga Rp.75.000, Voucher Game AXIS! 325 UC PUBG Mobile. Berlaku untuk 1 akun.', '2022-03-03 14:54:08', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `product_types`
--

CREATE TABLE `product_types` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `product_types`
--

INSERT INTO `product_types` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Pulsa', '2022-03-03 14:43:31', NULL),
(2, 'Paket Data', '2022-03-03 14:43:31', NULL),
(3, 'Voucher', '2022-03-03 14:43:31', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `payment_method_id` int(11) DEFAULT NULL,
  `status` varchar(10) DEFAULT NULL,
  `total_qty` int(11) DEFAULT 0,
  `total_price` decimal(25,2) DEFAULT 0.00,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 'Proses', 3, '49000.00', '2022-03-03 14:57:56', NULL),
(2, 1, 2, 'Batal', 3, '57000.00', '2022-03-03 14:57:56', NULL),
(3, 1, 3, 'Sukses', 3, '50000.00', '2022-03-03 14:57:56', NULL),
(4, 2, 1, 'Proses', 3, '88000.00', '2022-03-03 14:57:56', NULL),
(5, 2, 2, 'Sukses', 3, '118000.00', '2022-03-03 14:57:56', NULL),
(6, 2, 3, 'Sukses', 3, '178000.00', '2022-03-03 14:57:56', NULL),
(7, 3, 1, 'Proses', 3, '137000.00', '2022-03-03 14:57:56', NULL),
(8, 3, 2, 'Proses', 3, '109000.00', '2022-03-03 14:57:56', NULL),
(9, 3, 3, 'Sukses', 3, '49000.00', '2022-03-03 14:57:56', NULL),
(10, 4, 1, 'Sukses', 3, '57000.00', '2022-03-03 14:57:56', NULL),
(11, 4, 2, 'Batal', 3, '50000.00', '2022-03-03 14:57:56', NULL),
(12, 4, 3, 'Sukses', 3, '88000.00', '2022-03-03 14:57:56', NULL),
(13, 5, 1, 'Sukses', 3, '118000.00', '2022-03-03 14:57:56', NULL),
(14, 5, 2, 'Sukses', 3, '178000.00', '2022-03-03 14:57:56', NULL),
(15, 5, 3, 'Sukses', 3, '137000.00', '2022-03-03 14:57:56', NULL);

--
-- Triggers `transactions`
--
DELIMITER $$
CREATE TRIGGER `delete_data_transaction` AFTER DELETE ON `transactions` FOR EACH ROW BEGIN
DECLARE v_transaction_id INT(11);
SET v_transaction_id=OLD.id;
DELETE FROM transaction_details WHERE transaction_id=v_transaction_id;
END
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `transaction_details`
--

CREATE TABLE `transaction_details` (
  `transaction_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `status` varchar(10) DEFAULT NULL,
  `qty` int(11) DEFAULT 0,
  `price` decimal(25,2) DEFAULT 0.00,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transaction_details`
--

INSERT INTO `transaction_details` (`transaction_id`, `product_id`, `status`, `qty`, `price`, `created_at`, `updated_at`) VALUES
(1, 3, 'Proses', 1, '15000.00', '2022-03-03 14:59:40', NULL),
(2, 3, 'Batal', 1, '15000.00', '2022-03-03 14:59:40', NULL),
(2, 4, 'Batal', 1, '20000.00', '2022-03-03 14:59:40', NULL),
(3, 3, 'Sukses', 1, '15000.00', '2022-03-03 14:59:40', NULL),
(3, 4, 'Sukses', 1, '20000.00', '2022-03-03 14:59:40', NULL),
(3, 5, 'Sukses', 1, '15000.00', '2022-03-03 14:59:40', NULL),
(4, 4, 'Proses', 1, '20000.00', '2022-03-03 14:59:40', NULL),
(4, 5, 'Proses', 1, '15000.00', '2022-03-03 14:59:40', NULL),
(4, 6, 'Proses', 1, '53000.00', '2022-03-03 14:59:40', NULL),
(5, 5, 'Sukses', 1, '15000.00', '2022-03-03 14:59:40', NULL),
(5, 6, 'Sukses', 1, '53000.00', '2022-03-03 14:59:40', NULL),
(5, 7, 'Sukses', 1, '50000.00', '2022-03-03 14:59:40', NULL),
(6, 6, 'Sukses', 1, '53000.00', '2022-03-03 14:59:40', NULL),
(6, 7, 'Sukses', 1, '50000.00', '2022-03-03 14:59:40', NULL),
(6, 8, 'Sukses', 1, '75000.00', '2022-03-03 14:59:40', NULL),
(7, 7, 'Proses', 1, '50000.00', '2022-03-03 14:59:40', NULL),
(7, 8, 'Proses', 1, '75000.00', '2022-03-03 14:59:40', NULL),
(8, 8, 'Proses', 1, '75000.00', '2022-03-03 14:59:40', NULL),
(9, 3, 'Sukses', 1, '15000.00', '2022-03-03 14:59:40', NULL),
(10, 3, 'Sukses', 1, '15000.00', '2022-03-03 14:59:40', NULL),
(10, 4, 'Sukses', 1, '20000.00', '2022-03-03 14:59:40', NULL),
(11, 3, 'Batal', 1, '15000.00', '2022-03-03 14:59:40', NULL),
(11, 4, 'Batal', 1, '20000.00', '2022-03-03 14:59:40', NULL),
(11, 5, 'Batal', 1, '15000.00', '2022-03-03 14:59:40', NULL),
(12, 4, 'Sukses', 1, '20000.00', '2022-03-03 14:59:40', NULL),
(12, 5, 'Sukses', 1, '15000.00', '2022-03-03 14:59:40', NULL),
(12, 6, 'Sukses', 1, '53000.00', '2022-03-03 14:59:40', NULL),
(13, 5, 'Sukses', 1, '15000.00', '2022-03-03 14:59:40', NULL),
(13, 6, 'Sukses', 1, '53000.00', '2022-03-03 14:59:40', NULL),
(13, 7, 'Sukses', 1, '50000.00', '2022-03-03 14:59:40', NULL),
(14, 6, 'Sukses', 1, '53000.00', '2022-03-03 14:59:40', NULL),
(14, 7, 'Sukses', 1, '50000.00', '2022-03-03 14:59:40', NULL),
(14, 8, 'Sukses', 1, '75000.00', '2022-03-03 14:59:40', NULL),
(15, 7, 'Sukses', 1, '50000.00', '2022-03-03 14:59:40', NULL),
(15, 8, 'Sukses', 1, '75000.00', '2022-03-03 14:59:40', NULL);

--
-- Triggers `transaction_details`
--
DELIMITER $$
CREATE TRIGGER `update_data_transaction` AFTER DELETE ON `transaction_details` FOR EACH ROW BEGIN
DECLARE v_transaction_id INT(11);
SET v_transaction_id=OLD.transaction_id;
UPDATE transactions SET total_qty=(SELECT SUM(transaction_details.qty) FROM transaction_details WHERE transaction_details.transaction_id=v_transaction_id) WHERE id=v_transaction_id;
END
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `status` smallint(1) DEFAULT NULL,
  `dob` date DEFAULT NULL,
  `gender` char(1) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `status`, `dob`, `gender`, `created_at`, `updated_at`) VALUES
(1, 'Imanuel', 1, '2001-12-24', 'M', '2022-03-03 14:57:20', NULL),
(2, 'Tifanny', 1, '2003-12-11', 'F', '2022-03-03 14:57:20', NULL),
(3, 'Yanuar', 1, '2002-01-20', 'M', '2022-03-03 14:57:20', NULL),
(4, 'Grace', 1, '2000-04-15', 'F', '2022-03-03 14:57:20', NULL),
(5, 'Elton', 1, '2001-06-27', 'M', '2022-03-03 14:57:20', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `operators`
--
ALTER TABLE `operators`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `payment_methods`
--
ALTER TABLE `payment_methods`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `product_type_id` (`product_type_id`),
  ADD KEY `operator_id` (`operator_id`);

--
-- Indexes for table `product_descriptions`
--
ALTER TABLE `product_descriptions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `product_types`
--
ALTER TABLE `product_types`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `payment_method_id` (`payment_method_id`);

--
-- Indexes for table `transaction_details`
--
ALTER TABLE `transaction_details`
  ADD PRIMARY KEY (`transaction_id`,`product_id`),
  ADD KEY `product_id` (`product_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `operators`
--
ALTER TABLE `operators`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `payment_methods`
--
ALTER TABLE `payment_methods`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `product_descriptions`
--
ALTER TABLE `product_descriptions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `product_types`
--
ALTER TABLE `product_types`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `products_ibfk_1` FOREIGN KEY (`product_type_id`) REFERENCES `product_types` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `products_ibfk_2` FOREIGN KEY (`operator_id`) REFERENCES `operators` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `product_descriptions`
--
ALTER TABLE `product_descriptions`
  ADD CONSTRAINT `product_descriptions_ibfk_1` FOREIGN KEY (`id`) REFERENCES `products` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `transactions`
--
ALTER TABLE `transactions`
  ADD CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  ADD CONSTRAINT `transactions_ibfk_2` FOREIGN KEY (`payment_method_id`) REFERENCES `payment_methods` (`id`) ON DELETE SET NULL ON UPDATE CASCADE;

--
-- Constraints for table `transaction_details`
--
ALTER TABLE `transaction_details`
  ADD CONSTRAINT `transaction_details_ibfk_1` FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `transaction_details_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
