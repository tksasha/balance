# frozen_string_literal: true

module Items
  class GetResultService < ApplicationService
    def initialize(action_name, params)
      @action_name = action_name

      @params = params
    end

    def call
      return ::Items::GetResourceService.call(@params) if edit?

      "::Items::#{ @action_name.camelize }Service".constantize.call(@params)
    end

    private

    def edit?
      @action_name == 'edit'
    end
  end
end
