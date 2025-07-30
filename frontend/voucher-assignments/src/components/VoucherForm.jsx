import React, { useState } from "react";
import "../VoucherForm.css";

const VoucherForm = () => {
  const basePath = "http://127.0.0.1:5000/api";
  const [formData, setFormData] = useState({
    name: "",
    id: "",
    flightNumber: "",
    date: "",
    aircraft: "ATR",
  });

  const [seats, setSeats] = useState([]);
  const [message, setMessage] = useState("");
  const controller = new AbortController();
  const timeoutId = setTimeout(() => controller.abort(), 5000);

  const allFieldsFilled = formData.name.trim() !== "" && formData.id.trim() !== "" && formData.flightNumber.trim() !== "" && formData.date.trim() !== "" && formData.aircraft.trim() !== "";

  const handleChange = (e) => {
    setFormData((prev) => ({
      ...prev,
      [e.target.name]: e.target.value,
    }));
  };

  const handleGenerate = async () => {
    setMessage("🔎 Checking for existing vouchers...");

    try {
      const checkRes = await fetch(basePath + "/check", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          flightNumber: formData.flightNumber,
          date: formData.date,
        }),
      });
      const checkData = await checkRes.json();

      if (checkData.data.exists == true) {
        setMessage("❌ Vouchers already exist for this flight.");
        setSeats([]);
        return;
      }
    } catch (error) {
      if (error.name === "AbortError") {
        setMessage("⏰ Request timed out.");
      } else {
        console.log("err: ", error);
        setMessage("❌ An error occurred.");
      }
    } finally {
      clearTimeout(timeoutId);
    }

    setMessage("✨ Generating vouchers...");

    try {
      const genRes = await fetch(basePath + "/generate", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(formData),
        signal: controller.signal,
      });
      const genData = await genRes.json();

      if (genData.success) {
        setSeats(genData.data.seats);
        setMessage("✅ Vouchers generated successfully!");
      } else {
        setMessage("❌ " + genData.message);
      }
    } catch (error) {
      if (error.name === "AbortError") {
        setMessage("⏰ Request timed out.");
      } else {
        setMessage("❌ An error occurred.");
      }
    } finally {
      clearTimeout(timeoutId);
    }
  };

  return (
    <div className="voucher-container">
      <h2>✈️ Crew Voucher Generator</h2>
      <div className="form-group">
        <label>Crew Name</label>
        <input name="name" onChange={handleChange} />
      </div>
      <div className="form-group">
        <label>Crew ID</label>
        <input name="id" onChange={handleChange} />
      </div>
      <div className="form-group">
        <label>Flight Number</label>
        <input name="flightNumber" onChange={handleChange} />
      </div>
      <div className="form-group">
        <label>Flight Date</label>
        <input name="date" type="date" pattern="\d{4}-\d{2}-\d{2}" onChange={handleChange} />
      </div>
      <div className="form-group">
        <label>Aircraft Type</label>
        <select name="aircraft" onChange={handleChange}>
          <option value="ATR">ATR</option>
          <option value="Airbus 320">Airbus 320</option>
          <option value="Boeing 737 Max">Boeing 737 Max</option>
        </select>
      </div>
      <button className="generate-button" disabled={!allFieldsFilled} onClick={handleGenerate}>
        🎫 Generate Vouchers
      </button>
      <p className="message">{message}</p>
      {seats.length > 0 && (
        <div className="seat-list">
          <h4>🪑 Assigned Seats</h4>
          <ul>
            {seats.map((seat) => (
              <li key={seat}>{seat}</li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
};

export default VoucherForm;
