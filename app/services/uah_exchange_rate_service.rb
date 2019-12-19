# frozen_string_literal: true

class UAHExchangeRateService
  attr_reader :date

  def initialize(date = Date.today)
    @date = date
  end

  def save
    create_usd

    create_rub
  end

  private

  def rates
    @rates ||= NBUExchangeRateService.rates date
  end

  def create_usd
    ExchangeRate.create from: :uah, to: :usd, date: date, rate: rates[:usd]
  end

  def create_rub
    ExchangeRate.create from: :uah, to: :rub, date: date, rate: rates[:rub]
  end

  class << self
    def create(*args)
      new(*args).save
    end
  end
end
