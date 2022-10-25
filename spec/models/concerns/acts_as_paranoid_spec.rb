# frozen_string_literal: true

RSpec.describe ActsAsParanoid do
  let(:described_class) do
    Class.new(ApplicationRecord) do
      include ActsAsParanoid

      self.table_name = :items
    end
  end

  describe '#destroy' do
    before do
      allow(subject).to receive(:touch).with(:deleted_at)

      subject.destroy
    end

    it { is_expected.to have_received(:touch).with(:deleted_at) }
  end

  describe '.default_scope' do
    subject { described_class.all }

    let(:sql) { described_class.unscoped.where(deleted_at: nil).to_sql }

    its(:to_sql) { is_expected.to eq sql }
  end
end
