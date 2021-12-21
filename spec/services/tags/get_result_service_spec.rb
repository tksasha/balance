# frozen_string_literal: true

RSpec.describe Tags::GetResultService do
  subject { described_class.new action_name, category, params }

  let(:category) { build :category }

  let(:params) { double }

  describe '#call' do
    context do
      let(:action_name) { 'new' }

      before { allow(Tags::InitializeService).to receive(:call).with(category).and_return(:result) }

      its(:call) { should eq :result }
    end

    context do
      let(:action_name) { 'create' }

      before { allow(Tags::CreateService).to receive(:call).with(category, params).and_return(:result) }

      its(:call) { should eq :result }
    end
  end
end
