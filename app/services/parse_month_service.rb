# frozen_string_literal: true

class ParseMonthService
  attr_reader :month

  def initialize(params)
    self.month = params[:month]
  end

  def call
    month
  end

  private

  def month=(object)
    @month = \
      begin
        Month.parse(object)
      rescue ArgumentError
        Month.now
      end
  end

  class << self
    def call(*args)
      new(*args).call
    end
  end
end
