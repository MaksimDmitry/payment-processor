const fs = require('fs');
const { parse } = require('csv-parse');

module.exports = class PaymentParser {
  constructor(filePath) {
    this.filePath = filePath;
  }

  parse() {
    return new Promise((resolve, reject) => {
      fs.createReadStream(this.filePath)
        .pipe(parse({ delimiter: ',', from: 'header' }))
        .on('data', (row) => {
          const payment = {
            id: row[0],
            amount: parseFloat(row[1]),
            currency: row[2],
            date: new Date(row[3]),
          };
          // Add more properties as needed
          resolve(payment);
        })
        .on('end', () => {
          resolve();
        })
        .on('error', (error) => {
          reject(error);
        });
    });
  }
};