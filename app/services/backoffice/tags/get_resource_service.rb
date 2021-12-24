# frozen_string_literal: true

module Backoffice
  module Tags
    class GetResourceService < ApplicationService
      def initialize(params)
        @id = params[:id]
      end

      def call
        Success.new(tag)
      end

      private

      def tag
        @tag ||= Tag.find(@id)
      end
    end
  end
end
