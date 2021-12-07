# frozen_string_literal: true

module Categories
  class InitializeService < ApplicationService
    def call
      Success.new(category)
    end

    private

    def category
      Category.new
    end
  end
end
