# frozen_string_literal: true

module Backoffice
  module Tags
    class GetResultService < ApplicationService
      def initialize(action_name, category, params)
        @action_name = action_name

        @category = category

        @params = params
      end

      def call
        return ::Backoffice::Tags::InitializeService.call(@category) if new?

        "::Backoffice::Tags::#{ @action_name.camelize }Service".constantize.call(@category, @params)
      end

      private

      def new?
        @action_name == 'new'
      end
    end
  end
end
