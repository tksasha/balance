# frozen_string_literal: true

require 'open-uri'

class NbuExchangeRateService
  def initialize(date = Time.zone.today)
    @date = date
  end

  # rubocop:disable Metrics/AbcSize
  def rates
    %w[USD RUB].each_with_object({}) do |currency, memo|
      doc.css('#exchangeRates tbody tr').each do |node|
        next unless node.css('td[data-label="Код літерний"]').text == currency

        units = node.css('td[data-label="Кількість одиниць"]').text.to_i

        rate = node.css('td[data-label="Курс"]').text.tr(',', '.').to_d

        memo[currency.downcase.to_sym] = rate / units
      end
    end
  end
  # rubocop:enable Metrics/AbcSize

  private

  def date
    @date.strftime '%d.%m.%Y'
  end

  def url
    "https://www.bank.gov.ua/markets/exchangerates?date=#{ date }&period=daily"
  end

  def doc
    @doc ||= Nokogiri::HTML URI.parse(url).open.read
  # TODO: spec me
  rescue Errno::ECONNRESET
    sleep 1

    retry
  end

  class << self
    def rates(*args)
      new(*args).rates
    end
  end
end
