import winston from "winston";
import chalk from "chalk";
import dotenv from "dotenv";

dotenv.config();

export const logger = winston.createLogger({
  level: process.env.DEBUG === "true" ? "debug" : "info",
  format: winston.format.combine(
    winston.format.timestamp(),
    winston.format.printf(({ level, message, timestamp }) => {
      let levelColor = "";
      switch (level) {
        case "error":
          levelColor = chalk.red(level);
          break;
        case "warn":
          levelColor = chalk.yellow(level);
          break;
        case "info":
          levelColor = chalk.green(level);
          break;
        case "debug":
          levelColor = chalk.blue(level);
          break;
        default:
          levelColor = level;
      }
      return `${chalk.gray(timestamp)} :: ${levelColor} -> ${message}`;
    })
  ),
  transports: [new winston.transports.Console()],
});
