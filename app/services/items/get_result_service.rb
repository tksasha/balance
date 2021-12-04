# frozen_string_literal: true

module Items
  class GetResultService < ApplicationService
    def initialize(action_name, params)
      @action_name = action_name

      @params = params
    end

    def call
      "::Items::#{ @action_name.camelize }Service".constantize.call(@params)
    end
  end
end
