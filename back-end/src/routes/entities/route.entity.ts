import { Prop, raw, Schema, SchemaFactory } from '@nestjs/mongoose';

// Cria um tipo "RouteDocument" que é a junção desses dois que está recebendo
export type RouteDocument = Route & Document;

@Schema()
export class Route {
  @Prop()
  _id: string;

  @Prop()
  title: string;

  @Prop(
    // esse raw define o tipo do campo
    raw({
      lat: { type: Number },
      lng: { type: Number },
    }),
  )
  startPosition: { lat: number; lng: number };

  @Prop(
    raw({
      lat: { type: Number },
      lng: { type: Number },
    }),
  )
  endPosition: { lat: number; lng: number };
}

export const RouteSchema = SchemaFactory.createForClass(Route);
