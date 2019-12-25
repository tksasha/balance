# frozen_string_literal: true

module Backoffice
  class ExchangeRatesController < ApplicationController
    private

    def collection
      @collection ||= ExchangeRate.order(date: :desc).page(params[:page])
    end
  end
end
